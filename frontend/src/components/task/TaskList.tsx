import { useDeleteTask, useListTasks, useUpdateTask } from "@/api/generated/taskManagerApis";
import type { ModelTask } from "@/api/models/modelTask";
import {
  Loader2,
  Pencil,
  Trash2,
  ClipboardList,
  Check,
} from "lucide-react";
import { toast } from "sonner";
import { cn } from "@/lib/utils";

interface TaskListProps {
  onEdit: (task: ModelTask) => void;
}

export const TaskList = ({ onEdit }: TaskListProps) => {
  const { data: tasks, isLoading, refetch } = useListTasks();
  const { mutateAsync: deleteTask } = useDeleteTask();
  const { mutateAsync: updateTask } = useUpdateTask();

  const toggleStatus = async (task: ModelTask) => {
    try {
      const newStatus = task.status === "Pending" ? "Completed" : "Pending";
      await updateTask({
        id: task.id as string,
        data: { ...task, status: newStatus },
      });
      toast.success("Status updated", {
        description: `Task marked as ${newStatus.toLowerCase()}.`,
      });
      refetch();
    } catch (error) {
      toast.error("Error", {
        description: "Failed to update task status.",
      });
    }
  };

  const handleDelete = async (id: string) => {
    try {
      await deleteTask({ id });
      toast.success("Task deleted", {
        description: "The task has been deleted successfully.",
      });
      refetch();
    } catch (error) {
      toast.error("Error", {
        description: "Failed to delete task.",
      });
    }
  };

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[400px]">
        <Loader2 className="h-10 w-10 animate-spin text-indigo-600" />
      </div>
    );
  }

  if (!tasks?.data?.length) {
    return (
      <div className="flex flex-col items-center justify-center min-h-[400px] rounded-xl p-8">
        <ClipboardList className="h-20 w-20 text-indigo-200 mb-6" />
        <h3 className="text-2xl font-semibold text-gray-800 mb-3">
          Your task list is empty
        </h3>
        <p className="text-gray-600 max-w-md text-center">
          Create your first task to start organizing your work more efficiently.
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-3">
      {tasks.data.map((task: ModelTask) => (
        <div
          key={task.id}
          className="bg-white rounded-xl border border-gray-100 hover:border-indigo-100 transition-all duration-200 p-4 sm:p-5"
        >
          <div className="flex flex-col sm:flex-row items-start sm:items-center gap-4">

            <div className="flex-1 flex items-center gap-4">
              <label className="relative inline-flex items-center cursor-pointer">
                <input
                  type="checkbox"
                  className="sr-only"
                  checked={task.status === "Completed"}
                  onChange={() => toggleStatus(task)}
                />
                <div className={cn(
                  "w-5 h-5 flex items-center justify-center transition-all duration-200 ease-in-out border-2",
                  task.status === "Completed"
                    ? "bg-indigo-600 border-indigo-700"
                    : "bg-gray-100 border-gray-300"
                )}>
                  {task.status === "Completed" && (
                    <Check className="w-5 h-5 text-white" />
                  )}
                </div>

              </label>
              <div className="flex-1">
                <span
                  className={cn(
                    "text-lg font-medium",
                    task.status === "Completed"
                      ? "text-gray-400 line-through"
                      : "text-gray-800"
                  )}
                >
                  {task.name}
                </span>
                {task.description && (
                  <p className="text-sm text-gray-500 mt-1">{task.description}</p>
                )}
              </div>
            </div>
            <div className="flex items-center gap-2 self-end sm:self-center">
              <button
                onClick={() => onEdit(task)}
                className="p-2 text-gray-400 hover:text-indigo-600 hover:bg-indigo-50 rounded-lg transition-all duration-200"
              >
                <Pencil className="w-5 h-5" />
              </button>
              <button
                onClick={() => handleDelete(task.id as string)}
                className="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-all duration-200"
              >
                <Trash2 className="w-5 h-5" />
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};

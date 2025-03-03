import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { ModelTask } from "@/api/models/modelTask";
import { TaskDialog } from "./TaskForm";
import { TaskList } from "./TaskList";

export const TaskManager = () => {
  const [editingTask, setEditingTask] = useState<ModelTask | null>(null);
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleEdit = (task: ModelTask) => {
    setEditingTask(task);
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setEditingTask(null);
    setIsDialogOpen(false);
  };

  return (
    <div className="max-w-4xl mx-auto px-4 py-8 sm:px-6 lg:px-8">
      <div className="space-y-8">
        <header className="flex flex-col sm:flex-row justify-between items-center gap-6 pb-6 border-b border-gray-100">
          <div>
            <h1 className="text-3xl sm:text-4xl font-bold text-gray-900 mb-2">
              Task Flow
            </h1>
            <p className="text-gray-500">Manage your tasks efficiently</p>
          </div>
          <Button
            size="lg"
            className="w-full sm:w-auto rounded-xl bg-indigo-600 hover:bg-indigo-700 text-white shadow-sm hover:shadow-md transition-all duration-200"
            onClick={() => setIsDialogOpen(true)}
          >
            <Plus className="h-5 w-5 mr-2" />
            Create New Task
          </Button>
        </header>

        <main>
          <TaskList onEdit={handleEdit} />
        </main>

        <TaskDialog
          open={isDialogOpen}
          onOpenChange={handleDialogClose}
          editTask={editingTask}
          onSuccess={handleDialogClose}
        />
      </div>
    </div>
  );
};

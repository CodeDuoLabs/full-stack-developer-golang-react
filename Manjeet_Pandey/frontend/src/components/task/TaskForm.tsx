import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import type { ModelTask } from "@/api/models/modelTask";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { queryClient } from "@/main";
import { useCreateTask, useUpdateTask } from "@/api/generated/taskManagerApis";

const taskSchema = z.object({
  name: z.string().min(1, "Task name is required").max(100),
  description: z.string()
});

type TaskFormValues = z.infer<typeof taskSchema>;

interface TaskDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  editTask?: ModelTask | null;
  onSuccess?: () => void;
}

export const TaskDialog = ({
  open,
  onOpenChange,
  editTask,
  onSuccess,
}: TaskDialogProps) => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { mutateAsync: createTask } = useCreateTask();
  const { mutateAsync: updateTask } = useUpdateTask();

  const form = useForm<TaskFormValues>({
    resolver: zodResolver(taskSchema),
    defaultValues: {
      name: editTask?.name || "",
      description: editTask?.description || "",
    },
  });

  useEffect(() => {
    if (!editTask) return;
    form.reset({
      name: editTask.name,
      description: editTask.description
    });
  }, [editTask, form]);

  const onSubmit = async (data: TaskFormValues) => {
    try {
      setIsSubmitting(true);
      if (editTask?.id) {
        await updateTask({
          id: editTask.id,
          data: { ...data, status: editTask.status },
        });
        toast.success("Task updated", {
          description: "Your task has been updated successfully.",
        });
      } else {
        await createTask({ data: { ...data, status: "Pending" } });
        toast.success("Task created", {
          description: "Your new task has been created successfully.",
        });
      }
      queryClient.invalidateQueries({ queryKey: ["/tasks"] });
      form.reset();
      onSuccess?.();
    } catch (error) {
      toast.error("Error", {
        description: "Something went wrong. Please try again.",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px] bg-white rounded-lg shadow-lg border-0">
        <DialogHeader className="border-b pb-4">
          <div className="flex justify-between items-center">
            <DialogTitle className="text-xl font-semibold text-gray-900">
              {editTask ? "Edit Task" : "Create New Task"}
            </DialogTitle>
          </div>
        </DialogHeader>
        <div className="py-4">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Task Name</FormLabel>
                    <FormControl>
                      <Input placeholder="Enter task name..." {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>

                )}
              />
              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Description</FormLabel>
                    <FormControl>
                      <Input placeholder="Enter task Description..." {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>

                )}
              />
              <Button
                type="submit"
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white"
                disabled={isSubmitting}
              >
                {isSubmitting && (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                )}
                {editTask ? "Update Task" : "Create Task"}
              </Button>
            </form>
          </Form>
        </div>
      </DialogContent>
    </Dialog>
  );
};

import { useState } from "react";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { usePostTasks } from "@/api/generated/taskManagerApis";

interface TaskFormProps {
  onTaskAdded: () => void;
}

export const TaskForm = ({ onTaskAdded }: TaskFormProps) => {
  const [name, setName] = useState("");
  const { mutateAsync: createTask } = usePostTasks();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;
    await createTask({
      data: {
        name,
        status: "Pending"
      }
    })
    setName("");
    onTaskAdded();
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Input
        placeholder="Task Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        className="w-full"
      />
      <Button type="submit" className="w-full">
        Add Task
      </Button>
    </form>
  );
};

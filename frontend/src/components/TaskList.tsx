import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../components/ui/table";
import { Button } from "../components/ui/button";
import { useDeleteTasksId, useGetTasks, usePutTasksId } from "@/api/generated/taskManagerApis";
import { ModelTask } from "@/api/models/modelTask";

export const TaskList = () => {
  // const { getTasks, deleteTask, updateTask } = useApi();

  const { data: tasks } = useGetTasks()

  const { mutateAsync: deleteTask } = useDeleteTasksId()

  const { mutateAsync: updateTask } = usePutTasksId()



  const toggleStatus = async (task: ModelTask) => {
    const newStatus = task.status === "Pending" ? "Completed" : "Pending";
    await updateTask({ id: task.id as string, data: { status: newStatus } })
  };

  const handleDelete = async (id: string) => {
    await deleteTask({id: id});
  };

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Status</TableHead>
          <TableHead>Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {(tasks?.data.data || [] ).map((task: ModelTask) => (
          <TableRow key={task.id}>
            <TableCell>{task.name}</TableCell>
            <TableCell>{task.status}</TableCell>
            <TableCell className="space-x-2">
              <Button variant="outline" onClick={() => toggleStatus(task)}>
                Toggle
              </Button>
              <Button variant="destructive" onClick={() => handleDelete(task.id || '')}>
                Delete
              </Button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

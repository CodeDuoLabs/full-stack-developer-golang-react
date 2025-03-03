import { Toaster } from "sonner";
import { TaskManager } from "./components/task";

function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-rose-100 to-teal-100 p-4 md:p-8">
      <div className="max-w-5xl mx-auto">
        <TaskManager />
      </div>
      <Toaster position="top-center" expand={true} richColors />
    </div>
  );
}

export default App;

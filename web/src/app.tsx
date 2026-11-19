import { Toaster } from "./components/ui/sonner";
import AppRoutes from "./routes/app-routes";

const App = () => {
  return (
    <>
      <AppRoutes />
      <Toaster richColors position="top-right" />
    </>
  );
};

export default App;

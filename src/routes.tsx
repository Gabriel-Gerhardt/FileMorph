import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App";
import Any from "./pages/Any";
import Home from "./pages/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      { path: "", element: <Home /> },
      { path: "any", element: <Any /> },
    ],
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}

export default Routes;

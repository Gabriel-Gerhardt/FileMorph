import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App";
import Any from "./pages/Any";
import Home from "./pages/Home";
import Convert from "./pages/Convert";
import Login from "./pages/Login";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      { path: "", element: <Home /> },
      { path: "any", element: <Any /> },
      { path: "convert", element: <Convert /> },
      { path: "login", element: <Login /> },
    ],
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}

export default Routes;

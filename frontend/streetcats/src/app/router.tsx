import { createBrowserRouter } from "react-router-dom"
import { AppLayout } from "../components/layout/AppLayout"
import { HomePage } from "../pages/HomaPage"

export const router = createBrowserRouter([
    {
        path: "/",
        element: <AppLayout />,
        children: [
            {
                index: true,
                element: <HomePage />,
            },
        ],
    },
])
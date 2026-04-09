import { createBrowserRouter, Navigate } from "react-router-dom"
import { AppLayout } from "../components/layout/AppLayout"
import { HomePage } from "../pages/HomaPage"
import { LoginPage } from "../pages/LoginPage"
import { RegisterPage } from "../pages/RegisterPage"

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
    {
        path: "/login",
        element: <LoginPage />,
    },
    {
        path: "/register",
        element: <RegisterPage />,
    },
    {
        path: "*",
        element: <Navigate to="/" replace />,
    }
])
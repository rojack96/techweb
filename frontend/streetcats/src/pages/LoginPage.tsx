

import { useState } from "react"
import { useNavigate, Link } from "react-router-dom"
import { Form, Input, Button, message, Card } from "antd"
import { httpClient } from "../services/httpClient"

export function LoginPage() {
    const [loading, setLoading] = useState(false)
    const navigate = useNavigate()

    const onFinish = async (values: { username: string; password: string }) => {
        setLoading(true)
        try {
            const data = await httpClient<{ username?: string }>("/auth/login", {
                method: "POST",
                body: JSON.stringify(values),
            })

            const username = data.username ?? values.username
            localStorage.setItem("username", username)

            message.success("Login effettuato con successo!")
            navigate("/", { replace: true })
        } catch (error) {
            message.error("Credenziali non valide. Riprova.")
        } finally {
            setLoading(false)
        }
    }

    return (
        <div style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            minHeight: '100vh',
            background: '#f0f2f5'
        }}>
            <Card title="Accedi" style={{ width: 400 }}>
                <Form
                    name="login"
                    onFinish={onFinish}
                    autoComplete="off"
                    layout="vertical"
                >
                    <Form.Item
                        label="Username"
                        name="username"
                        rules={[{ required: true, message: 'Inserisci il tuo username!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Password"
                        name="password"
                        rules={[{ required: true, message: 'Inserisci la tua password!' }]}
                    >
                        <Input.Password />
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit" loading={loading} block>
                            Accedi
                        </Button>
                    </Form.Item>

                    <Form.Item style={{ textAlign: 'center' }}>
                        Non hai un account? <Link to="/register">Registrati</Link>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    )
}
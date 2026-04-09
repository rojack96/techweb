

import { useState } from "react"
import { Form, Input, Button, message, Card } from "antd"
import { Link } from "react-router-dom"

export function LoginPage() {
    const [loading, setLoading] = useState(false)

    const onFinish = async (values: { username: string; password: string }) => {
        setLoading(true)
        try {
            // TODO: Sostituire con chiamata API reale
            const response = await fetch('/api/auth/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(values),
            })

            if (response.ok) {
                const data = await response.json()
                // TODO: Gestire il token o redirect
                message.success('Login effettuato con successo!')
            } else {
                message.error('Credenziali non valide. Riprova.')
            }
        } catch (error) {
            message.error('Errore durante il login. Riprova più tardi.')
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
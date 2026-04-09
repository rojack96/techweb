import { useState } from "react"
import { Form, Input, Button, Select, message, Card } from "antd"
import { Link } from "react-router-dom"

// TODO: Nel futuro, sostituire con chiamata API per ottenere le regioni
const regioniItaliane = [
    "Abruzzo",
    "Basilicata",
    "Calabria",
    "Campania",
    "Emilia-Romagna",
    "Friuli-Venezia Giulia",
    "Lazio",
    "Liguria",
    "Lombardia",
    "Marche",
    "Molise",
    "Piemonte",
    "Puglia",
    "Sardegna",
    "Sicilia",
    "Toscana",
    "Trentino-Alto Adige",
    "Umbria",
    "Valle d'Aosta",
    "Veneto"
]

export function RegisterPage() {
    const [loading, setLoading] = useState(false)

    const onFinish = async (values: {
        username: string;
        password: string;
        nome: string;
        cognome: string;
        regione: string;
    }) => {
        setLoading(true)
        try {
            // TODO: Sostituire con chiamata API reale
            const response = await fetch('/api/auth/register', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(values),
            })

            if (response.ok) {
                const data = await response.json()
                message.success('Registrazione effettuata con successo!')
                // TODO: Redirect a login o dashboard
            } else {
                message.error('Errore durante la registrazione. Riprova.')
            }
        } catch (error) {
            message.error('Errore durante la registrazione. Riprova più tardi.')
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
            <Card title="Registrati" style={{ width: 400 }}>
                <Form
                    name="register"
                    onFinish={onFinish}
                    autoComplete="off"
                    layout="vertical"
                >
                    <Form.Item
                        label="Username"
                        name="username"
                        rules={[{ required: true, message: 'Inserisci un username!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Password"
                        name="password"
                        rules={[{ required: true, message: 'Inserisci una password!' }]}
                    >
                        <Input.Password />
                    </Form.Item>

                    <Form.Item
                        label="Nome"
                        name="nome"
                        rules={[{ required: true, message: 'Inserisci il tuo nome!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Cognome"
                        name="cognome"
                        rules={[{ required: true, message: 'Inserisci il tuo cognome!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Regione"
                        name="regione"
                        rules={[{ required: true, message: 'Seleziona la tua regione!' }]}
                    >
                        <Select
                            placeholder="Seleziona la tua regione"
                            options={regioniItaliane.map((regione) => ({
                                label: regione,
                                value: regione
                            }))}>
                        </Select>
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit" loading={loading} block>
                            Registrati
                        </Button>
                    </Form.Item>

                    <Form.Item style={{ textAlign: 'center' }}>
                        Hai già un account? <Link to="/login">Accedi</Link>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    )
}
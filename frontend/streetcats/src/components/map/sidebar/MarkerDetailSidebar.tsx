import { useState } from "react"
import { Card, Typography, Button, Input, Divider } from "antd"
import { XMarkdown } from "@ant-design/x-markdown"

const { Title, Text } = Typography
const { TextArea } = Input

interface MarkerDetailSidebarProps {
    isOpen: boolean
    marker: {
        title: string
        description: string
        breed: string
        id: string
    } | null
    onClose: () => void
}

export function MarkerDetailSidebar({ isOpen, marker, onClose }: MarkerDetailSidebarProps) {
    // Simula lo stato di login - in futuro sarà collegato all'auth reale
    const [isLoggedIn] = useState(true) // Cambia a false per testare modalità non loggato
    const [comments, setComments] = useState<string[]>([
        "Ottimo avvistamento! Ho visto un gatto simile nella zona.",
        "Grazie per la segnalazione, controllerò domani.",
        "Il gatto sembra in buona salute, complimenti per la descrizione dettagliata!"
    ])
    const [newComment, setNewComment] = useState("")

    if (!isOpen || !marker) return null

    const handleAddComment = () => {
        if (newComment.trim()) {
            setComments(prev => [...prev, newComment.trim()])
            setNewComment("")
        }
    }

    return (
        <aside
            style={{
                width: 500,
                borderLeft: "1px solid #f0f0f0",
                background: "#fff",
                padding: 16,
                boxSizing: "border-box",
                overflowY: "auto",
                position: "absolute",
                right: 0,
                top: 0,
                bottom: 0,
                zIndex: 1000
            }}
        >
            <Card
                size="small"
                title={
                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                        <Title level={5} style={{ margin: 0 }}>Dettagli Marker</Title>
                        <button
                            onClick={onClose}
                            style={{ background: "none", border: "none", cursor: "pointer", fontSize: 18 }}
                        >
                            ✕
                        </button>
                    </div>
                }
            >
                {/* Titolo del marker */}
                <div style={{ marginBottom: 16 }}>
                    <Title level={4} style={{ margin: 0 }}>{marker.title}</Title>
                    <Text type="secondary">{marker.breed}</Text>
                </div>

                <Divider />

                {/* Descrizione completa */}
                <div style={{ marginBottom: 24 }}>
                    <Title level={5} style={{ marginBottom: 12 }}>Descrizione</Title>
                    <div style={{ fontSize: 14, lineHeight: 1.6 }}>
                        <XMarkdown>{marker.description}</XMarkdown>
                    </div>
                </div>

                <Divider />

                {/* Sezione commenti */}
                <div>
                    <Title level={5} style={{ marginBottom: 16 }}>Commenti ({comments.length})</Title>

                    {/* Lista commenti esistenti */}
                    <div style={{ marginBottom: 16, maxHeight: 200, overflowY: "auto" }}>
                        {comments.length === 0 ? (
                            <Text type="secondary">Nessun commento ancora.</Text>
                        ) : (
                            comments.map((comment, index) => (
                                <div
                                    key={index}
                                    style={{
                                        padding: 12,
                                        background: "#f5f5f5",
                                        borderRadius: 4,
                                        marginBottom: 8,
                                        fontSize: 14
                                    }}
                                >
                                    <XMarkdown>{comment}</XMarkdown>
                                </div>
                            ))
                        )}
                    </div>

                    {/* Form per nuovo commento - solo se loggato */}
                    {isLoggedIn ? (
                        <div>
                            <TextArea
                                value={newComment}
                                onChange={(e) => setNewComment(e.target.value)}
                                placeholder="Scrivi un commento... (supporta **grassetto**, *corsivo*, ecc.)"
                                rows={3}
                                style={{ marginBottom: 8 }}
                            />
                            <Button
                                type="primary"
                                onClick={handleAddComment}
                                disabled={!newComment.trim()}
                                size="small"
                            >
                                Aggiungi Commento
                            </Button>
                        </div>
                    ) : (
                        <div style={{ padding: 12, background: "#fff2f0", border: "1px solid #ffccc7", borderRadius: 4 }}>
                            <Text type="secondary">
                                Devi essere loggato per aggiungere commenti.
                            </Text>
                        </div>
                    )}
                </div>
            </Card>
        </aside>
    )
}
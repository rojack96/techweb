import { Card, Select, Input, Typography, Button, Space } from "antd"
import { XMarkdown } from "@ant-design/x-markdown"
import { breedOptions } from "../../hooks/useMarkerForm"

const { TextArea } = Input
const { Title } = Typography

interface AddMarkerSidebarProps {
    isOpen: boolean
    selectedBreed: string | undefined
    notes: string
    onClose: () => void
    onBreedChange: (breed: string | undefined) => void
    onNotesChange: (notes: string) => void
    onSave: () => void
    onCancel: () => void
}

export function AddMarkerSidebar({
    isOpen,
    selectedBreed,
    notes,
    onClose,
    onBreedChange,
    onNotesChange,
    onSave,
    onCancel,
}: AddMarkerSidebarProps) {
    if (!isOpen) return null

    return (
        <aside
            style={{
                width: 450,
                borderLeft: "1px solid #f0f0f0",
                background: "#fff",
                padding: 16,
                boxSizing: "border-box",
                overflowY: "auto",
            }}
        >
            <Card
                size="small"
                title={<Title level={5}>Aggiungi Marker</Title>}
                extra={
                    <button
                        onClick={onClose}
                        style={{ background: "none", border: "none", cursor: "pointer", fontSize: 18 }}
                    >
                        ✕
                    </button>
                }
            >
                <div style={{ marginBottom: 16 }}>
                    <label style={{ display: "block", marginBottom: 8 }}>Razza</label>
                    <Select
                        value={selectedBreed}
                        onChange={(value) => onBreedChange(value)}
                        options={breedOptions}
                        placeholder="Seleziona razza"
                        style={{ width: "100%" }}
                        allowClear
                    />
                </div>
                <div style={{ marginBottom: 16 }}>
                    <label style={{ display: "block", marginBottom: 8 }}>Note (Markdown)</label>
                    <TextArea
                        value={notes}
                        onChange={(e) => onNotesChange(e.target.value)}
                        rows={5}
                        placeholder="Descrizione in markdown"
                    />
                    {notes && (
                        <div
                            style={{
                                marginTop: 12,
                                border: "1px solid #d9d9d9",
                                borderRadius: 4,
                                padding: 12,
                                background: "#fafafa",
                            }}
                        >
                            <small style={{ color: "#666", display: "block", marginBottom: 8 }}>
                                Anteprima:
                            </small>
                            <XMarkdown>{notes}</XMarkdown>
                        </div>
                    )}
                </div>
                <div>
                    <Space style={{ width: "100%" }}>
                        <Button type="primary" onClick={onSave} style={{ flex: 1 }}>
                            Salva
                        </Button>
                        <Button onClick={onCancel} style={{ flex: 1 }}>
                            Annulla
                        </Button>
                    </Space>
                </div>
            </Card>
        </aside>
    )
}

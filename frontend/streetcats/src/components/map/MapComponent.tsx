import { useState } from "react"
import { MapContainer, TileLayer } from "react-leaflet"
import { CustomMarkers } from "./marker/Marker"
import "leaflet/dist/leaflet.css"
import MarkerClusterGroup from "react-leaflet-cluster"
import { LocationMarker } from "./marker/LocationMarker"
import { InsertMarker } from "./button/InsertMarker"
import { Card, Select, Input, Typography, Button, Space } from "antd"
import { XMarkdown } from "@ant-design/x-markdown"

const { TextArea } = Input
const { Title } = Typography

export function MapComponent() {
    const [markerMode, setMarkerMode] = useState(false)
    const [sidebarOpen, setSidebarOpen] = useState(false)
    const [extraMarkers, setExtraMarkers] = useState<{ position: [number, number]; markerText: string; id: string }[]>([])
    const [pendingMarker, setPendingMarker] = useState<{ position: [number, number]; id: string } | null>(null)
    const [selectedBreed, setSelectedBreed] = useState<string | undefined>(undefined)
    const [notes, setNotes] = useState("")

    const breedOptions = [
        { value: "siamese", label: "Siamese" },
        { value: "maine_coon", label: "Maine Coon" },
        { value: "persian", label: "Persiano" },
        { value: "ragdoll", label: "Ragdoll" },
        { value: "sphynx", label: "Sphynx" },
    ]

    const handleMarkerPlaced = (position: [number, number]) => {
        const id = Date.now().toString()
        setExtraMarkers(prev => [...prev, { position, markerText: "Modifica...", id }])
        setPendingMarker({ position, id })
        setSidebarOpen(true)
    }

    const handleSave = () => {
        if (!pendingMarker) return

        const breedText = selectedBreed ? breedOptions.find(b => b.value === selectedBreed)?.label ?? selectedBreed : "Razza non selezionata"
        const markerText = `${breedText}${notes ? ` - ${notes}` : ""}`

        setExtraMarkers(prev =>
            prev.map(marker =>
                marker.id === pendingMarker.id
                    ? { ...marker, markerText }
                    : marker
            )
        )
        resetForm()
    }

    const handleCancel = () => {
        if (pendingMarker !== null) {
            setExtraMarkers(prev => prev.filter(marker => marker.id !== pendingMarker.id))
        }
        resetForm()
    }

    const resetForm = () => {
        setPendingMarker(null)
        setSidebarOpen(false)
        setMarkerMode(false)
        setSelectedBreed(undefined)
        setNotes("")
    }

    const closeSidebar = () => {
        setSidebarOpen(false)
        setMarkerMode(false)
    }

    return (
        <div style={{ display: "flex", height: "100%", width: "100%" }}>
            <div style={{ position: "relative", flex: sidebarOpen ? "0 0 calc(100% - 450px)" : 1, transition: "flex 0.2s" }}>
                <InsertMarker active={markerMode} onToggle={() => setMarkerMode(prev => !prev)} />
                <MapContainer
                    center={[42.5, 12.5]}
                    zoom={6}
                    scrollWheelZoom={true}
                    style={{ height: "100%", width: "100%" }}
                >
                    {/* Satellite layer*/}
                    <TileLayer
                        attribution='&copy; Esri'
                        url="https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}"
                    />

                    {/* Labels layer */}
                    <TileLayer
                        attribution='&copy; Esri Map data © Esri — Sources: Esri, Maxar, Earthstar Geographics, CNES/Airbus DS, USDA, USGS, AeroGRID, IGN, Garmin, HERE, OpenStreetMap contributors, and the GIS User Community.'
                        url="https://services.arcgisonline.com/ArcGIS/rest/services/Reference/World_Transportation/MapServer/tile/{z}/{y}/{x}"
                    />
                    <LocationMarker active={markerMode} onMarkerPlaced={handleMarkerPlaced} />

                    <MarkerClusterGroup>
                        <CustomMarkers markers={[
                            { position: [40.851112, 14.268901], markerText: "Cat sighting 1" },
                            { position: [40.843567, 14.255432], markerText: "Cat sighting 2" },
                            { position: [40.862345, 14.272198], markerText: "Cat sighting 3" },
                            { position: [40.836789, 14.243567], markerText: "Cat sighting 4" },
                            { position: [40.857901, 14.289345], markerText: "Cat sighting 5" },
                            { position: [40.869234, 14.261987], markerText: "Cat sighting 6" },
                            { position: [40.828765, 14.235678], markerText: "Cat sighting 7" },
                            { position: [40.874512, 14.277654], markerText: "Cat sighting 8" },
                            { position: [40.845678, 14.298765], markerText: "Cat sighting 9" },
                            { position: [40.833456, 14.259876], markerText: "Cat sighting 10" },
                            ...extraMarkers
                        ]} />
                    </MarkerClusterGroup>
                </MapContainer>
            </div>

            {sidebarOpen && (
                <aside style={{
                    width: 450,
                    borderLeft: "1px solid #f0f0f0",
                    background: "#fff",
                    padding: 16,
                    boxSizing: "border-box",
                    overflowY: "auto"
                }}>
                    <Card size="small" title={<Title level={5}>Aggiungi Marker</Title>} extra={<button onClick={closeSidebar} style={{ background: "none", border: "none", cursor: "pointer", fontSize: 18 }}>✕</button>}>
                        <div style={{ marginBottom: 16 }}>
                            <label style={{ display: "block", marginBottom: 8 }}>Razza</label>
                            <Select
                                value={selectedBreed}
                                onChange={(value) => setSelectedBreed(value)}
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
                                onChange={(e) => setNotes(e.target.value)}
                                rows={5}
                                placeholder="Descrizione in markdown"
                            />
                            {notes && (
                                <div style={{ marginTop: 12, border: "1px solid #d9d9d9", borderRadius: 4, padding: 12, background: "#fafafa" }}>
                                    <small style={{ color: "#666", display: "block", marginBottom: 8 }}>Anteprima:</small>
                                    <XMarkdown>{notes}</XMarkdown>
                                </div>
                            )}
                        </div>
                        <div>
                            <Space style={{ width: "100%" }}>
                                <Button type="primary" onClick={handleSave} style={{ flex: 1 }}>
                                    Salva
                                </Button>
                                <Button onClick={handleCancel} style={{ flex: 1 }}>
                                    Annulla
                                </Button>
                            </Space>
                        </div>
                    </Card>
                </aside>
            )}
        </div>
    )
}
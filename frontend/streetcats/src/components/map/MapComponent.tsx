import { MapContainer, TileLayer } from "react-leaflet"
import { CustomMarkers } from "./marker/Marker"
import "leaflet/dist/leaflet.css"
import MarkerClusterGroup from "react-leaflet-cluster"
import { LocationMarker } from "./marker/LocationMarker"
import { InsertMarker } from "./button/InsertMarker"
import { AddMarkerSidebar } from "./sidebar/AddMarkerSidebar"
import { useMarkerForm } from "../../hooks/useMarkerForm"

export function MapComponent() {
    const {
        markerMode,
        sidebarOpen,
        extraMarkers,
        selectedBreed,
        markerTitle,
        markerDescription,
        setMarkerMode,
        setSelectedBreed,
        setMarkerTitle,
        setMarkerDescription,
        handleMarkerPlaced,
        handleSave,
        handleCancel,
        closeSidebar,
    } = useMarkerForm()

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
                            ...extraMarkers.map(m => {
                                // Limita la descrizione a 2-3 righe nel tooltip
                                const descLines = m.description.split('\n').slice(0, 3).join('\n')
                                return {
                                    position: m.position,
                                    markerText: `# ${m.title}\n\n**${m.breed}**\n\n${descLines}`
                                }
                            })
                        ]} />
                    </MarkerClusterGroup>
                </MapContainer>
            </div>

            <AddMarkerSidebar
                isOpen={sidebarOpen}
                selectedBreed={selectedBreed}
                markerTitle={markerTitle}
                markerDescription={markerDescription}
                onClose={closeSidebar}
                onBreedChange={setSelectedBreed}
                onTitleChange={setMarkerTitle}
                onDescriptionChange={setMarkerDescription}
                onSave={handleSave}
                onCancel={handleCancel}
            />
        </div>
    )
}
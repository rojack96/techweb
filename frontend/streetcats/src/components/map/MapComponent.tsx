import { useState } from "react"
import { MapContainer, TileLayer } from "react-leaflet"
import { CustomMarkers } from "./marker/Marker"
import "leaflet/dist/leaflet.css"
import MarkerClusterGroup from "react-leaflet-cluster"
import { LocationMarker } from "./marker/LocationMarker"
import { InsertMarker } from "./button/InsertMarker"

export function MapComponent() {
    const [markerMode, setMarkerMode] = useState(false)
    const [extraMarkers, setExtraMarkers] = useState<{ position: [number, number]; markerText: string }[]>([])

    const handleMarkerPlaced = (position: [number, number]) => {
        setExtraMarkers(prev => [...prev, { position, markerText: "Nuovo gattino" }])
        setMarkerMode(false)
    }

    return (
        <div style={{ height: "100%", width: "100%", position: "relative" }}>
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
    )
}
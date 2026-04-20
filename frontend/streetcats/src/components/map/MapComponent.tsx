import { MapContainer, TileLayer } from "react-leaflet"
import { CustomMarkers } from "./marker/Marker"
import "leaflet/dist/leaflet.css"
import MarkerClusterGroup from "react-leaflet-cluster"
import { LocationMarker } from "./marker/LocationMarker"
import { InsertMarker } from "./button/InsertMarker"
import { AddMarkerSidebar } from "./sidebar/AddMarkerSidebar"
import { MarkerDetailSidebar } from "./sidebar/MarkerDetailSidebar"
import { useMarkerForm } from "../../hooks/useMarkerForm"
import { useSightings } from "../../features/sightings/hooks"

export function MapComponent() {
    const {
        markerMode,
        sidebarOpen,
        // TODO capire gli extraMakers come lavorano
        extraMarkers,
        selectedBreed,
        markerTitle,
        markerDescription,
        detailSidebarOpen,
        selectedMarkerForDetail,
        breedOptions,
        setMarkerMode,
        setSelectedBreed,
        setMarkerTitle,
        setMarkerDescription,
        handleMarkerPlaced,
        handleSave,
        handleCancel,
        closeSidebar,
        handleExpandMarker,
        closeDetailSidebar,
    } = useMarkerForm()

    const { data: sightings, isLoading, error } = useSightings()
    console.log("Sightings:", sightings, "Loading:", isLoading, "Error:", error)
    const sightingMarkers = sightings?.response?.map(sighting => ({
        position: sighting.position,
        title: sighting.title,
        description: sighting.description,
        breed: sighting.breed,
        id: sighting.id,
    })) ?? []

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
                        <CustomMarkers
                            markers={[...sightingMarkers, ...extraMarkers]}
                            onExpand={handleExpandMarker}
                        />
                    </MarkerClusterGroup>
                </MapContainer>
            </div>

            <AddMarkerSidebar
                isOpen={sidebarOpen}
                selectedBreed={selectedBreed}
                markerTitle={markerTitle}
                markerDescription={markerDescription}
                breedOptions={breedOptions}
                onClose={closeSidebar}
                onBreedChange={setSelectedBreed}
                onTitleChange={setMarkerTitle}
                onDescriptionChange={setMarkerDescription}
                onSave={handleSave}
                onCancel={handleCancel}
            />

            <MarkerDetailSidebar
                isOpen={detailSidebarOpen}
                marker={selectedMarkerForDetail}
                onClose={closeDetailSidebar}
            />
        </div>
    )
}
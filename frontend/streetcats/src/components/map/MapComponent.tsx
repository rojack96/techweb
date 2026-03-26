import { MapContainer, TileLayer } from "react-leaflet"
import { CustomMarkers } from "./marker/Marker"
import "leaflet/dist/leaflet.css"
import MarkerClusterGroup from "react-leaflet-cluster";

export function MapComponent() {
    return (
        <MapContainer
            center={[40.82922739413729, 14.189834511099555]}
            zoom={27}
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

            <MarkerClusterGroup>
                <CustomMarkers positions={[
                    [40.851112, 14.268901],
                    [40.843567, 14.255432],
                    [40.862345, 14.272198],
                    [40.836789, 14.243567],
                    [40.857901, 14.289345],
                    [40.869234, 14.261987],
                    [40.828765, 14.235678],
                    [40.874512, 14.277654],
                    [40.845678, 14.298765],
                    [40.833456, 14.259876],
                ]} />
            </MarkerClusterGroup>
        </MapContainer>
    )
}
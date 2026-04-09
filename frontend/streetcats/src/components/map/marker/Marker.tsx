import L from "leaflet"
import { Marker, Popup } from "react-leaflet"
import { XMarkdown } from '@ant-design/x-markdown';
import catMarkerUrl from "../../../assets/cat-svgrepo-com.svg"

interface CustomMarkersProps {
    position: [number, number]
    markerText: string
}

const customIcon = L.icon({
    iconUrl: catMarkerUrl,
    iconSize: [35, 35],
    iconAnchor: [18, 2],
    className: ""
});


interface CatMarkerProps {
    position: [number, number]
    markerText: string
}

export function CatMarker({ position, markerText }: CatMarkerProps) {
    return (
        <Marker
            icon={customIcon}
            position={position}
        >
            <Popup maxWidth={300} minWidth={250}>
                <div style={{ fontSize: 14, lineHeight: 1.6, padding: 8 }}>
                    <XMarkdown>{markerText}</XMarkdown>
                </div>
            </Popup>
        </Marker>
    )
}

export function CustomMarkers({ markers }: { markers: CustomMarkersProps[] }) {
    return (
        <>
            {markers.map((marker, index) => (
                <CatMarker
                    key={index}
                    position={marker.position}
                    markerText={marker.markerText}
                />
            ))}
        </>
    )
}
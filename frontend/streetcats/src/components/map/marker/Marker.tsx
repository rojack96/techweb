import L from "leaflet"
import { Marker, Popup } from "react-leaflet"
import { XMarkdown } from '@ant-design/x-markdown';
import catMarkerUrl from "../../../assets/cat-svgrepo-com.svg"

interface CustomMarkersProps {
    position: [number, number]
    markerText: string
    title?: string
    description?: string
    breed?: string
    id?: string
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
    title?: string
    description?: string
    breed?: string
    id?: string
    onExpand?: (marker: { title: string; description: string; breed: string; id: string }) => void
}

export function CatMarker({ position, markerText, title, description, breed, id, onExpand }: CatMarkerProps) {
    const hasFullData = title && description && breed && id
    const descLines = description ? description.split('\n').slice(0, 2).join('\n') : ''
    const hasMoreContent = description && description.split('\n').length > 2

    return (
        <Marker
            icon={customIcon}
            position={position}
        >
            <Popup maxWidth={300} minWidth={250}>
                <div style={{ fontSize: 14, lineHeight: 1.6, padding: 8 }}>
                    {hasFullData ? (
                        <>
                            <div style={{ fontSize: 16, fontWeight: 'bold', marginBottom: 8 }}>{title}</div>
                            <div style={{ fontWeight: 'bold', color: '#666', marginBottom: 8 }}>{breed}</div>
                            <div style={{ marginBottom: hasMoreContent ? 12 : 0 }}>
                                <XMarkdown>{descLines}</XMarkdown>
                            </div>
                            {hasMoreContent && (
                                <div style={{ marginTop: 12, textAlign: 'center' }}>
                                    <button
                                        onClick={() => onExpand?.({ title, description, breed, id })}
                                        style={{
                                            background: '#1890ff',
                                            color: 'white',
                                            border: 'none',
                                            borderRadius: 4,
                                            padding: '4px 12px',
                                            cursor: 'pointer',
                                            fontSize: 12
                                        }}
                                    >
                                        Leggi di più
                                    </button>
                                </div>
                            )}
                        </>
                    ) : (
                        <XMarkdown>{markerText}</XMarkdown>
                    )}
                </div>
            </Popup>
        </Marker>
    )
}

export function CustomMarkers({ markers, onExpand }: { markers: CustomMarkersProps[], onExpand?: (marker: { title: string; description: string; breed: string; id: string }) => void }) {
    return (
        <>
            {markers.map((marker, index) => (
                <CatMarker
                    key={index}
                    position={marker.position}
                    markerText={marker.markerText}
                    title={marker.title}
                    description={marker.description}
                    breed={marker.breed}
                    id={marker.id}
                    onExpand={onExpand}
                />
            ))}
        </>
    )
}
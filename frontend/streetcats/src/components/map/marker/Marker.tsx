import L from "leaflet"
import { Marker, Popup } from "react-leaflet"
import { XMarkdown } from '@ant-design/x-markdown';
import catMarkerUrl from "../../../assets/cat-svgrepo-com.svg"

interface CustomMarkersProps {
    id: string
    position: [number, number]
    title: string
    description: string
    breed: string
}

const customIcon = L.icon({
    iconUrl: catMarkerUrl,
    iconSize: [35, 35],
    iconAnchor: [18, 2],
    className: ""
});


interface CatMarkerProps {
    position: [number, number]
    title: string
    description: string
    breed: string
    id: string
    onExpand?: (marker: { title: string; description: string; breed: string; id: string }) => void
}

export function CatMarker({ position, title, description, breed, id, onExpand }: CatMarkerProps) {
    const descLines = description ? description.split('\n').slice(0, 2).join('\n') : ''
    const shouldTruncate = description && description.split('\n').length > 2

    return (
        <Marker
            icon={customIcon}
            position={position}
        >
            <Popup maxWidth={300} minWidth={250}>
                <div style={{ fontSize: 14, lineHeight: 1.6, padding: 8 }}>
                    <>
                        <div style={{ fontSize: 16, fontWeight: 'bold', marginBottom: 8 }}>{title}</div>
                        <div style={{ fontWeight: 'bold', color: '#666', marginBottom: 8 }}>{breed}</div>
                        <div style={{ marginBottom: 12 }}>
                            <XMarkdown>{shouldTruncate ? descLines : description}</XMarkdown>
                            {shouldTruncate && (
                                <div style={{ fontSize: 12, color: '#999', marginTop: 4 }}>
                                    (continua...)
                                </div>
                            )}
                        </div>
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
                    </>
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
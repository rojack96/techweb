import { useMapEvents } from "react-leaflet"
import { useState } from "react"
import type { LatLng } from 'leaflet'
import { CatMarker } from "./Marker"

export function LocationMarker() {
    const [position, setPosition] = useState<LatLng | null>(null)

    useMapEvents({
        click(e) {
            console.log("Clicked at: ", e.latlng)
            setPosition(e.latlng)
        },
    })


    return position ?
        <CatMarker
            key=""
            position={[position.lat, position.lng]}
            markerText="Your Marker Text" />
        : null
}
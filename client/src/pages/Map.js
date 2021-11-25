import React from 'react';

import 'leaflet/dist/leaflet.css';

import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import { Polygon } from 'react-leaflet';

const position = [48.3766022,5.028867]

const coords = [
    {lat: 48.3766022, lng: 5.028867},
    {lat: 48.3766031, lng: 5.0288412},
    {lat: 48.3765823, lng: 5.0288367},
    {lat: 48.3765804, lng: 5.0288629},
    {lat: 48.3766022, lng: 5.028867}
]

function Map() {
    return (
        <MapContainer center={position} zoom={13} scrollWheelZoom={false}>
            <TileLayer
            attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            <Marker position={position}>
            <Popup>
                A pretty CSS3 popup. <br /> Easily customizable.
            </Popup>
            </Marker>
            <Polygon color={"red"} positions={coords} />
    </MapContainer>
    );
}

export default Map;
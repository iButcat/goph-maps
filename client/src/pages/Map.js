import React, { useState, useEffect } from 'react';

import Leaflet from "leaflet";
import 'leaflet/dist/leaflet.css';
import axios from 'axios';

import { MapContainer, TileLayer, Marker, Popup, Polyline } from 'react-leaflet';

import iconShadow from "leaflet/dist/images/marker-shadow.png";

let DefaultIcon = Leaflet.icon({
  iconSize: [25, 41],
  iconAnchor: [2, 41],
  popupAnchor: [2, -40],
  iconUrl: "https://unpkg.com/leaflet@1.6/dist/images/marker-icon.png",
  shadowUrl: iconShadow
});

Leaflet.Marker.prototype.options.icon = DefaultIcon;

const position = [49.2625,5.9619]

function Map() {
    const [coordinates, setCoordinates] = useState([]);

    const getCoordinates = () => {
        return axios.get("http://localhost:8080/hello");
    };

    useEffect(() => {
        getCoordinates().then((response) => {
            if (response.status === 200) {
                setCoordinates(response.data.graph);
                return;
            } else {
                return;
            }
        }).catch((error) => {console.log(error)});
    },[]);

    return (
        <MapContainer center={position} zoom={13} scrollWheelZoom={false}>
            <TileLayer
            attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {coordinates.length !== 0 && coordinates.map((coordinate) => (
                <Marker 
                position={[coordinate.Point.geometry[1], coordinate.Point.geometry[0]]}
                icon={DefaultIcon}>
                <Popup>
                   {coordinate.Point.name}
                </Popup>
                </Marker>
            ))}
    </MapContainer>
    );
}

export default Map;
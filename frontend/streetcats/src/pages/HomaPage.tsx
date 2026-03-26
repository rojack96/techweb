import { Row, Col } from "antd"
import { MapComponent } from "../components/map/MapComponent"

export function HomePage() {
    return (
        <Row gutter={[16, 16]} justify="center" align="top">
            <Col xs={24} md={6} lg={6}>
                <div style={{ minHeight: "520px", maxHeight: "600px", border: "1px solid #f0f0f0", padding: 12 }}>
                    Lista sinistra
                </div>
            </Col>

            <Col xs={24} md={12} lg={12}>
                <div style={{
                    height: "700px",
                    width: "100%",
                    border: "1px solid #f0f0f0",
                    overflow: "hidden",
                }}>
                    <MapComponent />
                </div>
            </Col>

            <Col xs={24} md={6} lg={6}>
                <div style={{ minHeight: "80vh", border: "1px solid #f0f0f0", padding: 12 }}>
                    Lista destra
                </div>
            </Col>
        </Row >
    )
}
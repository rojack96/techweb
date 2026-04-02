import { Row, Col } from "antd"
import { MapComponent } from "../components/map/MapComponent"

export function HomePage() {
    return (
        <div>
            <Row gutter={[8, 16]} justify="start" align="top">
                <Col xs={24} md={12} lg={24}>
                    <div style={{
                        height: "90vh",
                        border: "1px solid #f0f0f0",
                        overflow: "hidden",
                        marginRight: 10,
                        marginLeft: 10,
                    }}>
                        <MapComponent />
                    </div>
                </Col>
            </Row >
        </div >
    )
}
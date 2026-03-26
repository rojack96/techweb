import { useState } from "react"
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from "@ant-design/icons"
import { Button, Layout, Menu, theme } from "antd"
import { Outlet } from "react-router-dom"

const { Header, Sider, Content } = Layout

// AppLayout è il layout principale dell'applicazione. 
// Contiene una sidebar con un menu e un header con un pulsante per espandere o collassare la sidebar. 
// Il contenuto principale viene renderizzato all'interno del componente Outlet, che è un placeholder per i componenti figli definiti nelle rotte.
export function AppLayout() {
    const [collapsed, setCollapsed] = useState(true)

    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken()

    return (
        <Layout>
            <Sider trigger={null} collapsible collapsed={collapsed}>
                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={["1"]}
                    items={[
                        { key: "1", icon: <UserOutlined />, label: "nav 1" },
                        { key: "2", icon: <VideoCameraOutlined />, label: "nav 2" },
                        { key: "3", icon: <UploadOutlined />, label: "nav 3" },
                    ]}
                />
            </Sider>

            <Layout>
                <Header style={{ padding: 0, background: colorBgContainer }}>
                    <Button
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                        onClick={() => setCollapsed(!collapsed)}
                    />
                </Header>

                <Content
                    style={{
                        margin: "24px 16px",
                        padding: 24,
                        background: colorBgContainer,
                        borderRadius: borderRadiusLG,
                    }}
                >
                    <Outlet />
                </Content>
            </Layout>
        </Layout>
    )
}
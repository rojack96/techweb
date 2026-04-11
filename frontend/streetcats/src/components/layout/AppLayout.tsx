import { useEffect, useState } from "react"
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from "@ant-design/icons"
import { Button, Dropdown, Layout, Menu, message, theme } from "antd"
import { Outlet, useNavigate } from "react-router-dom"
import { httpClient } from "../../services/httpClient"

const { Header, Sider, Content } = Layout

// AppLayout è il layout principale dell'applicazione. 
// Contiene una sidebar con un menu e un header con un pulsante per espandere o collassare la sidebar. 
// Il contenuto principale viene renderizzato all'interno del componente Outlet, che è un placeholder per i componenti figli definiti nelle rotte.
export function AppLayout() {
    const [collapsed, setCollapsed] = useState(true)
    const [isAdmin] = useState(false)
    const [user, setUser] = useState<{ username: string } | null>(null)
    const navigate = useNavigate()

    useEffect(() => {
        const loadUser = async () => {
            try {
                const data = await httpClient<{ username: string }>("/auth/me")
                if (data?.username) {
                    setUser({ username: data.username })
                }
            } catch {
                setUser(null)
            }
        }

        loadUser()
    }, [])

    const {
        token: { colorBgContainer },
    } = theme.useToken()

    return (
        <Layout style={{ minHeight: "100vh" }}>
            {isAdmin && (
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
            )}

            <Layout>
                <Header style={{ padding: 0, background: colorBgContainer }}>
                    <div style={{
                        display: 'flex', justifyContent: 'space-between', alignItems: 'center',
                        padding: '16px 16px', width: '100%', boxSizing: 'border-box'
                    }}>
                        <div style={{ display: 'flex', alignItems: 'center', minWidth: 48 }}>
                            {isAdmin && (
                                <Button
                                    type="text"
                                    icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                                    onClick={() => setCollapsed(!collapsed)}
                                />
                            )}
                        </div>
                        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-end', minWidth: 48 }}>
                            <Dropdown menu={{
                                items: user ? [
                                    { key: "info", label: `Ciao ${user.username}`, disabled: true },
                                    { type: "divider" },
                                    { key: "logout", label: "Logout" },
                                ] : [
                                    { key: "login", label: "Login" },
                                ], onClick: async ({ key }) => {
                                    if (key === "login") {
                                        navigate("/login")
                                        return
                                    }
                                    if (key === "logout") {
                                        try {
                                            await httpClient("/auth/logout", { method: "POST" })
                                            message.success("Logout effettuato")
                                        } catch {
                                            message.error("Errore durante il logout")
                                        } finally {
                                            localStorage.removeItem("username")
                                            setUser(null)
                                            navigate("/login")
                                        }
                                    }
                                }
                            }} trigger={["click"]} placement="bottomRight">
                                <Button type="text" icon={<UserOutlined />} />
                            </Dropdown>
                        </div>
                    </div>
                </Header>

                <Content
                    style={{
                        margin: "0",
                        padding: "0",
                        background: colorBgContainer,
                        minHeight: "calc(100vh - 64px)", // header altezza 64px se usi default
                    }}
                >
                    <Outlet />
                </Content>
            </Layout>
        </Layout>
    )
}
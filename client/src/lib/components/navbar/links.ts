const links = [
    { name: "Agents", url: "/agents" },
    { name: "Favourites", url: "/favourites" },
    { name: "Login", url: "/auth/login" },
    { name: "Register", url: "/auth/register" },
]
export default links
export type Link = typeof links[number]
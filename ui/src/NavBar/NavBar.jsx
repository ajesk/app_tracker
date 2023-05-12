import { Flex, Heading } from "@chakra-ui/react"
import { Link } from "react-router-dom"

const NavBar = () => {
    return (
        <Flex w="19em" h="100%" bg='aquamarine'>
            <Heading marginLeft="0em">
                Application Tracker
            </Heading>
            <Link to="/">Applications</Link>
            <Link to="/companies">Companies</Link>
        </Flex>
    )
}

export default NavBar;
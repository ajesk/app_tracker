import {
    Flex, Icon
} from "@chakra-ui/react"
import { Link } from "react-router-dom"

const LinkItems = [
    { name: 'Home', to: '/' },
    { name: 'Applications', to: '/applications' },
    { name: 'Companies', to: '/companies' },
];

const NavBar = () => {
    return (
        <>
            {
                LinkItems.map((link) => (
                    <Link key={link.name} to={link.to} style={{ textDecoration: 'none' }} _focus={{ boxShadow: 'none' }}>
                        <Flex
                            align="center"
                            p="4"
                            mx="1"
                            borderRadius="lg"
                            role="group"
                            cursor="pointer"
                            _hover={{
                                bg: 'cyan.400',
                                color: 'white',
                            }}
                        >
                            {link.name}
                            {link.icon && (
                                <Icon
                                    mr="4"
                                    fontSize="16"
                                    _groupHover={{
                                        color: 'white',
                                    }}
                                    as={link.icon}
                                />
                            )}
                        </Flex>
                    </Link>
                ))
            }
        </>
    )
}

export default NavBar;
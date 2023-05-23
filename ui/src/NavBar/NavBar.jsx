import {
    Flex, Box, useColorModeValue, Text
} from "@chakra-ui/react"
import NavItems from './NavItems';

const NavBar = () => {
    return (
        <Box minH="100vh" bg={useColorModeValue('blue.100', 'gray.900')}>
            <Box
                bg={useColorModeValue('blue.50', 'gray.900')}
                w={{ base: 'full', md: 60 }}
                h="full"
                margin="0.125em"
                padding="0.25em"
            >
                <Flex
                    paddingTop="1em"
                    paddingBottom="1em"
                    h="20" 
                    alignItems="center" 
                    mx="11" 
                    justifyContent="space-between">
                    <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
                        Application Tracker
                    </Text>
                </Flex>
                <NavItems />
            </Box>
        </Box>
    )
}

export default NavBar;
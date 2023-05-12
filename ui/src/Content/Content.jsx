import { Flex, Text } from "@chakra-ui/react"
import { Route, Routes } from "react-router-dom";
import Applications from './Applications/Applications'
import Companies from './Companies/Companies'

const Content = () => {

    return (
        <Flex direction="column">
            <Routes>
                <Route path="/" exact element={<Applications />} />
                <Route path="/companies" element={<Companies />} />
            </Routes>
        </Flex>
    );
}

export default Content;
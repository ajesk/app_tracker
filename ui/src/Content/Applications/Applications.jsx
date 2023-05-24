import { Box, Button } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { getUserApplications } from "../../api/applictions";
import ApplicationTable from "./ApplicationTable";
import ApplicationModal from "./ApplicationModal";

const Applications = () => {
    const [applications, setApplications] = useState([]);
    const [isOpen, showModal] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await getUserApplications(1);
                setApplications(res);
            } catch (error) {
                console.error('Error:', error);
            }
        }
        fetchData();
    }, []);

    return (
        <Box>
            <Button onClick={() => showModal(true)}>Add Application</Button>
            <ApplicationTable applications={applications} />
            <ApplicationModal isOpen={isOpen} handleClose={() => showModal(false)} />
        </Box>
    );
};

export default Applications;
import { Box, Button } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { getUserApplications } from "../../api/applictions";
import ApplicationTable from "./ApplicationTable";
import CreateApplicationModal from "./CreateApplicationModal";
import ApplicationDetailModal from "./ApplicationDetailModal";

const defaultApplication = {
    isOpen: false,
    applicationId: 0
};

const Applications = () => {
    const [applications, setApplications] = useState([]);
    const [isOpen, showModal] = useState(false);
    const [applicationModalDetail, setApplicationModalDetail] = useState(defaultApplication);
    const showApplicaitonModal = (applicationId) => setApplicationModalDetail({isOpen: true, applicationId: applicationId}); 
    const hideApplicationModal = () => setApplicationModalDetail(defaultApplication);

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
            <ApplicationTable applications={applications} showApplicationModal={showApplicaitonModal} />
            <CreateApplicationModal isOpen={isOpen} handleClose={() => showModal(false)} />
            <ApplicationDetailModal 
                {...applicationModalDetail}
                handleClose={() => hideApplicationModal()} />
        </Box>
    );
};

export default Applications;
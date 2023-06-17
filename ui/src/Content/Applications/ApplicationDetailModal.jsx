import {
    Modal, ModalBody, ModalContent,
    ModalOverlay, ModalHeader, ModalCloseButton, Heading, Text, Flex, Box
} from "@chakra-ui/react";
import { getApplicationById } from "../../api/applictions";
import { useEffect, useState } from "react";

const ApplicationDetailModal = ({ handleClose, isOpen, applicationId }) => {
    const [application, setApplication] = useState({});

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await getApplicationById(applicationId);
                setApplication(res);
            } catch (error) {
                console.error('Error:', error);
            }
        }
        if (applicationId !== 0) fetchData();
    }, [applicationId]);

    const statuses = [
        {
            type: "applied",
            note: "",
            timestamp: "1-2-3"
        },
        {
            type: "screening",
            note: "typical screening",
            timestamp: "2-3-4"
        },
        {
            type: "technical interview",
            note: "panel 3 levels",
            timestamp: "3-4-5"
        },

    ]
    console.log(application)
    return (
        <Modal isOpen={isOpen} onClose={handleClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>{`${application.title} @ ${application.company}`}</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <Flex direction="row">
                        <Box flexGrow="1">
                            <Heading paddingBottom="1em" size="md">Details</Heading>
                            <Heading size="xs">Salary Range</ Heading>
                            <Text>
                                {`${application.min_salary} - ${application.max_salary}`}
                            </Text>
                            <br />
                            <Heading size="xs">Notes</ Heading>
                            <Text>
                                {application.notes}
                            </Text>
                            <br />
                            <Heading size="xs">Applied On</ Heading>
                            <Text>
                                {application.created_date}
                            </Text>
                        </Box>
                        <Box>
                            <Heading paddingBottom="1em" size="md">Status Updates</Heading>
                            {
                                statuses.map(status => 
                                    <Box rounded="base" margin="0.25em" padding="0.5em" border="1px" borderColor="gray.400">
                                        <Heading size="xs">{status.type}</ Heading>
                                        <Text>{status.note}</ Text>
                                        <Text fontSize="xs" color="gray.400">{status.timestamp}</ Text>
                                    </Box>
                                )
                            }
                        </Box>
                    </Flex>
                </ModalBody>
            </ModalContent>
        </Modal>
    )
};

export default ApplicationDetailModal;
import {
    Modal, ModalBody, ModalContent,
    ModalOverlay, ModalHeader, ModalCloseButton,
} from "@chakra-ui/react";
import CreateApplicationForm from "./CreateApplicationForm";

const ApplicationModal = ({ isOpen, handleClose }) => {
    const handleSubmit = (e, formData) => {
        fetch('API_URL', {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then((response) => response.json())
            .then((data) => {
                // Handle API response
                console.log(data);
                handleClose();
            })
            .catch((error) => {
                // Handle error
                console.error(error);
            });
    };

    return (
        <Modal isOpen={isOpen} onClose={handleClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>Create Application</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <CreateApplicationForm handleSubmit={handleSubmit} />
                </ModalBody>
            </ModalContent>
        </Modal>
    )
};

export default ApplicationModal;
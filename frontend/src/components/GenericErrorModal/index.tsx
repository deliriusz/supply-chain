import React from 'react'
import { Button, Modal, ModalProps } from 'semantic-ui-react'

function GenericErrorModal(props: ModalProps) {
   const [open, setOpen] = React.useState(props.open)

   return (
      <Modal
         onClose={() => setOpen(false)}
         onOpen={() => setOpen(true)}
         open={open}
      >
         <Modal.Header>{props.header}</Modal.Header>
         <Modal.Content>
            <Modal.Description>
               {props.description}
            </Modal.Description>
         </Modal.Content>
         <Modal.Actions>
            <Button
               content="Understood"
               labelPosition='right'
               icon='checkmark'
               onClick={() => setOpen(false)}
               negative
            />
         </Modal.Actions>
      </Modal>
   )
}

export default GenericErrorModal

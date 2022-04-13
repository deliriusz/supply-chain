import React, { useState } from "react";
import { Button, ButtonProps, Icon } from "semantic-ui-react";
import { processPayment } from "../../services/PaymentService";
import GenericModal from "../GenericErrorModal"

export type PurchaseButtonProps = ButtonProps & {
   purchaseProductId: number,
   purchaseAmount: number,
   purchaseAmountUnit?: string,
}

const PurchaseButton = (props: PurchaseButtonProps) => {
   const [paymentStatus, setPaymentStatus] = useState({ completed: false, error: false, message: "" })

   const handlePayment = (productId: number, amount: number) => {
      processPayment(productId, amount)
         .then(isSuccess => {
            const message = isSuccess ? "Payment successful" : "Payment could not be performed. Please try again later."
            setPaymentStatus({ completed: true, error: !isSuccess, message: message })
         }).catch(err => {
            setPaymentStatus({ completed: true, error: true, message: JSON.stringify(err) })
         })
   }

   return (
      <>
         <Button
            {...props}
            icon
            labelPosition="left"
            onClick={() => handlePayment(props.purchaseProductId, props.purchaseAmount)}
         >
            <Icon name='shopping cart' /> Buy
         </Button>
         <GenericModal
            header={paymentStatus.error ? "Payment not successful" : "Payment successful"}
            description={paymentStatus.message}
            open={paymentStatus.completed}
            isPositive={!paymentStatus.error}
            onClose={() => setPaymentStatus({ completed: false, error: false, message: "" })}
         />
      </>
   )
}

export default PurchaseButton;
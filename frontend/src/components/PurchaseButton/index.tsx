import React from "react";
import { Button, ButtonProps, Icon } from "semantic-ui-react";
import { processPayment } from "../../services/PaymentService";

export type PurchaseButtonProps = ButtonProps & {
   purchaseProductId: number,
   purchaseAmount: number,
   purchaseAmountUnit?: string,
}

const PurchaseButton = (props: PurchaseButtonProps) => {
   return (
      <Button
         {...props}
         icon
         labelPosition="left"
         onClick={() => processPayment(props.purchaseProductId, props.purchaseAmount)}
      >
         <Icon name='shopping cart' /> Buy
      </Button>
   )
}

export default PurchaseButton;
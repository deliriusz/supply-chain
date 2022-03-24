import React from "react";
import { GridColumn, Icon, Segment, Statistic } from "semantic-ui-react";
import SalesChart from "../../../components/SalesChart";

const StatisticsPane = () => {

   return (
      <>
         <GridColumn>
            <Segment compact>
               <Statistic.Group widths={'two'}>
                  <Statistic>
                     <Statistic.Value>22</Statistic.Value>
                     <Statistic.Label>Saves</Statistic.Label>
                  </Statistic>

                  <Statistic>
                     <Statistic.Value text>
                        Three
                        <br />
                        Thousand
                     </Statistic.Value>
                     <Statistic.Label>Signups</Statistic.Label>
                  </Statistic>

                  <Statistic>
                     <Statistic.Value>
                        <Icon name='plane' />5
                     </Statistic.Value>
                     <Statistic.Label>Flights</Statistic.Label>
                  </Statistic>

                  <Statistic>
                     <Statistic.Value>
                        42
                     </Statistic.Value>
                     <Statistic.Label>Team Members</Statistic.Label>
                  </Statistic>
               </Statistic.Group>
            </Segment>
            <SalesChart />
         </GridColumn>
      </>
   )
}

export default StatisticsPane;
import React, { Component } from 'react';
import PropTypes from 'prop-types';

import { ResponsiveLine } from '@nivo/line';

class TemperaturesChart extends Component {
  render() {
    const { data } = this.props;

    return (
      <ResponsiveLine
        data={data}
        margin={{
          top: 70,
          right: 50,
          bottom: 150,
          left: 70
        }}
        xScale={{
          type: 'time',
          format: '%Y-%m-%dT%H:%M:%S',
          precision: 'minute'
        }}
        stacked={true}
        curve="monotoneX"
        axisBottom={{
          orient: 'bottom',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: -90,
          legend: 'time',
          legendOffset: 110,
          legendPosition: 'center',
          format: '%Y-%m-%d %H:%M'
        }}
        axisLeft={{
          orient: 'left',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'temperature',
          legendOffset: -40,
          legendPosition: 'center'
        }}
        colors="set1"
        dotSize={10}
        dotColor="inherit:darker(0.3)"
        dotBorderWidth={2}
        dotBorderColor="#ffffff"
        enableDotLabel={true}
        dotLabel="y"
        dotLabelYOffset={-12}
        enableArea={true}
        animate={true}
        motionStiffness={90}
        motionDamping={15}
        legends={[
          {
            anchor: 'top-left',
            direction: 'row',
            justify: false,
            translateX: 0,
            translateY: -40,
            itemsSpacing: 0,
            itemDirection: 'left-to-right',
            itemWidth: 80,
            itemHeight: 20,
            itemOpacity: 0.75,
            symbolSize: 12,
            symbolShape: 'circle',
            symbolBorderColor: 'rgba(0, 0, 0, .5)',
            effects: [
              {
                on: 'hover',
                style: {
                  itemBackground: 'rgba(0, 0, 0, .03)',
                  itemOpacity: 1
                }
              }
            ]
          }
        ]}
      />
    );
  }
}

TemperaturesChart.propTypes = {
  data: PropTypes.array.isRequired
};

export default TemperaturesChart;

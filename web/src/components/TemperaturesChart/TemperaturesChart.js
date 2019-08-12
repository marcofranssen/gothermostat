import React, { Component } from 'react';
import PropTypes from 'prop-types';

import { ResponsiveLine } from '@nivo/line';

class TemperaturesChart extends Component {
  render() {
    const { data } = this.props;

    return (
      <ResponsiveLine
        data={data}
        margin={{ top: 70, right: 50, bottom: 150, left: 70 }}
        xScale={{
          type: 'time',
          format: '%Y-%m-%dT%H:%M:%S',
          precision: 'minute'
        }}
        xFormat="time:%d %b %H:%M"
        yScale={{ type: 'linear', stacked: false }}
        curve="monotoneX"
        axisBottom={{
          orient: 'bottom',
          tickValues: 'every 12 hours',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: -90,
          legend: 'time',
          legendOffset: -12,
          legendPosition: 'middle',
          format: '%d %b %H:%M'
        }}
        axisLeft={{
          orient: 'left',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'temperature',
          legendOffset: -40,
          legendPosition: 'middle'
        }}
        colors={{ scheme: 'pastel1' }}
        pointColor="white"
        pointSize={6}
        pointBorderWidth={2}
        pointBorderColor={{ from: 'serieColor', modifiers: [['darker', 0.3]] }}
        enablePointLabel={true}
        pointLabel="y"
        pointLabelYOffset={-12}
        enableArea={true}
        debugSlices={true}
        animate={true}
        motionStiffness={90}
        motionDamping={15}
        useMesh={true}
        legends={[
          {
            anchor: 'top-left',
            direction: 'column',
            justify: false,
            translateX: 0,
            translateY: -60,
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

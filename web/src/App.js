import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {
  withStyles,
  MuiThemeProvider,
  createMuiTheme
} from '@material-ui/core/styles';
import {
  AppBar,
  Toolbar,
  Box,
  CssBaseline,
  Typography
} from '@material-ui/core';
import { withFetching, TemperaturesChart } from './components';

const styles = theme => ({
  root: {
    flexGrow: 1
  },
  chart: {
    color: 'black',
    background: 'white',
    height: '650px'
  }
});

const theme = createMuiTheme({
  typography: {
    useNextVariants: true
  },
  palette: {
    type: 'dark'
  }
});

const safeLast = col => (col.length > 0 ? col[col.length - 1] : col[0]);

const mapNestData = type => t => ({
  id: t.name + ' ' + type,
  data: t.temperatures
    .map(temp => ({
      x: temp.timestamp.slice(0, 19), // strip milliseconds and timezone
      y: temp[type + 'TemperatureC']
    }))
    .reduce((acc, temp) => {
      let last = safeLast(acc);
      if (!last) {
        return [...acc, temp];
      } else {
        return last.y !== temp.y ? [...acc, temp] : acc;
      }
    }, [])
});

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { chartData: null };
  }

  componentDidUpdate() {
    const { data } = this.props;

    if (!data || this.state.chartData) {
      return;
    }
    const chartData = data.thermostats
      .map(mapNestData('ambient'))
      .concat(data.thermostats.map(mapNestData('target')));
    this.setState(prevState => ({ ...prevState, chartData }));
  }

  render() {
    const { classes, error, isLoading } = this.props;
    const { chartData } = this.state;

    return (
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        <div className={classes.root}>
          <AppBar position="static" color="default">
            <Toolbar>
              <Typography variant="h6" color="inherit">
                Go thermostat
              </Typography>
            </Toolbar>
          </AppBar>
          <Box className={classes.chart}>
            {!chartData || error || isLoading ? (
              'no data'
            ) : (
              <TemperaturesChart
                data={chartData}
                className={{ color: 'black' }}
              />
            )}
          </Box>
        </div>
      </MuiThemeProvider>
    );
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  isLoading: PropTypes.bool.isRequired,
  data: PropTypes.object,
  error: PropTypes.object
};

export default withFetching('/api/thermostat-data')(
  withStyles(styles, { withTheme: true })(App)
);

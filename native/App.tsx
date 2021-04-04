import React, { ReactElement } from 'react';
import { Provider } from 'react-redux';
import { ThemeProvider } from 'react-native-elements';
import { THEME } from '~~/constants/theme';
import { UiContext } from '~/lib/context';
import { createApplicationInitialState } from '~/lib/context/ui';
import store from '~/store';
import MainRoute from '~/routes';

const App = function App(): ReactElement {
  const [applicationState, setApplicationState] = React.useState(createApplicationInitialState());

  return (
    <Provider store={store}>
      <ThemeProvider theme={THEME}>
        <UiContext.Provider value={{ applicationState, setApplicationState }}>
          <MainRoute />
        </UiContext.Provider>
      </ThemeProvider>
    </Provider>
  );
};

export default App;
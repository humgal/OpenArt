// RootNavigation.js

import { createNavigationContainerRef } from '@react-navigation/native';

export const navigationRef = createNavigationContainerRef()

export  function navigate(name, params) {
  if (navigationRef.isReady()) {
    navigationRef.navigate(name, params);
  }else{
    console.log("navigate not ready")
  }
}

// add other navigation functions that you need and export them
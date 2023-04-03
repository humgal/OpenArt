import React from "react";
import { View } from "react-native";
import Head from "./src/components/common/Head/Head";
import BottomTab from "./src/components/common/BottomTab/BottomTab";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import Home from "./src/components/Home/Home";
import Creator from "./src/components/DiscoverCreator/DiscoverCreator";
import UploadArtWork from "./src/components/UploadArtWork/UploadArtWork";
import Profile from "./src/components/MyProfile/Profile";
import { navigationRef } from './src/components/common/RootNavigation';
import Menu from "./src/components/Menu/Menu";
import Search from "./src/components/Search/Search";

export default function App() {
  return (
    <View style={{ height: "100%" }}>
      <Head />
      <MyStack></MyStack>
      {/* <BottomTab></BottomTab> */}
    </View>
  );
}

const Tab = createBottomTabNavigator();
const Stack = createNativeStackNavigator();

const HomeTabs=()=>{
  return (
    <Tab.Navigator
    screenOptions={{headerShown: false,}}
    tabBar={props=><BottomTab {...props}/>}
      >
        <Tab.Screen name="Home" component={Home} />
        <Tab.Screen name="Creator" component={Creator} />
        <Tab.Screen name="UploadArtWork" component={UploadArtWork} />
        <Tab.Screen name="Profile" component={Profile} />
    </Tab.Navigator>
  );
}

const MyStack = () => {
  return (
    <NavigationContainer ref={navigationRef}>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
        }}>
        <Stack.Screen name="HomeTab" component={HomeTabs}/>
        <Stack.Screen  name="Menu" component={Menu}/>
        <Stack.Screen  name="Search" component={Search}/> 
      </Stack.Navigator>
    </NavigationContainer>
  );
};

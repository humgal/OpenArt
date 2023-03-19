import { StatusBar } from "expo-status-bar";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import Head from "./src/components/common/Head/Head";
import BottomTab,{navigationRef} from "./src/components/common/BottomTab/BottomTab";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Home from "./src/components/Home/Home";
import Creator from "./src/components/DiscoverCreator/Creator";
import UploadArtWork from "./src/components/UploadArtWork/UploadArtWork";
import Profile from "./src/components/MyProfile/Profile";

export default function App() {
  return (
    <View style={{ height: "100%" }}>
      <Head />
      <MyStack></MyStack>
      <BottomTab></BottomTab>
    </View>
  );
}

const Stack = createNativeStackNavigator();

const MyStack = () => {
  return (
    <NavigationContainer ref={navigationRef}>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
        }}
      >
        <Stack.Screen name="Home" component={Home} />
        <Stack.Screen name="Creator" component={Creator} />
        <Stack.Screen name="UploadArtWork" component={UploadArtWork} />
        <Stack.Screen name="Profile" component={Profile} />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

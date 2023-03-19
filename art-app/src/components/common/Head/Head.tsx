import React, { useState } from "react";
import { View, Image, StyleSheet } from "react-native"
import * as RootNavigation from "../RootNavigation";
import { IconButton } from "react-native-paper";
import Icon from "@expo/vector-icons/Ionicons";

const Head = () => {
  const [showMenuView, setMenuView] = useState(true);
  const [showNotificationView, setNotificationView] = useState(false);
  const [showProfileView, setProfileView] = useState(false);
  const [showCloseView, setCloseView] = useState(false);
  const [showSearchView, setSearchView] = useState(true);
  return (
    <View style={styles.container}>
      <Image
        style={styles.tinyLogo}
        source={require("../../../../assets/Logo.png")}
      />
      {showNotificationView && (
        <View style={styles.notification}>
                    <IconButton
            icon={props => <Icon name="notifications-outline" {...props} />}
            onPress={() => {
              
            }} 
          />
        </View>
      )}
      {showProfileView && (
        <View style={styles.profile}>
          <IconButton
            icon={props => <Icon name="md-person-outline" {...props} />}
            onPress={() => {
              
            }}
          />
        </View>
      )}
      {showSearchView && (
        <View style={styles.search}>
          <IconButton
            icon={props => <Icon name="search" {...props} />}
            onPress={() => {
                setCloseView(false);
                setMenuView(true);
                setSearchView(false);
                RootNavigation.navigate("Search");

            }}
          />
        </View>
      )}

      {showMenuView && (
        <View style={styles.menu}>
          <IconButton
            icon={props => <Icon name="menu" {...props} />}
            onPress={() => {
              setCloseView(true);
              setMenuView(false);
              setSearchView(true)
              RootNavigation.navigate("Menu");
            }}
          ></IconButton>
        </View>
      )}

      {showCloseView && (
        <View style={styles.close}>
          <IconButton
            icon={"close"}
            onPress={() => {
              setCloseView(false);
              setMenuView(true);
              setSearchView(true)
              if (RootNavigation.navigationRef.canGoBack()) {
                RootNavigation.navigate("Home");
              }
            }}
          ></IconButton>
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    marginTop: 50,
    flexDirection: "row",
    padding: 10,
  },
  tinyLogo: {
    width: 120,
    height: 30,
  },
  search: {
    position: "absolute",
    right: 50,

  },
  menu: {

    position: "absolute",
    right: 0,
  },
  notification: {

    position: "absolute",
    right: 90,
  },
  profile: {

    position: "absolute",
    right: 50,
  },
  close: {
    position: "absolute",
    right: 0,
  },
});

export default Head;

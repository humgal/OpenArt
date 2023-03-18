import React, { useState } from "react";
import { View, Image, StyleSheet } from "react-native";
import Svg, { Path, Circle } from "react-native-svg";

const Head = () => {
  const [showNotificationView, setShowNotificationView] = useState(false);
  const [showProfileView, setProfileView] = useState(false);
  return (
    <View style={styles.container}>
      <Image
        style={styles.tinyLogo}
        source={require("../../../../assets/Logo.png")}
      />
      {showNotificationView && (
        <View style={styles.notification}>
          <Svg width="26" height="24" viewBox="0 0 26 24" fill="none">
            <Path
              d="M20.3333 9.01449C20.3333 7.15413 19.5607 5.36997 18.1855 4.0545C16.8102 2.73902 14.9449 2 13 2C11.0551 2 9.18982 2.73902 7.81455 4.0545C6.43928 5.36997 5.66667 7.15413 5.66667 9.01449C5.66667 14.9484 3.21304 17.809 1.86406 18.9333C1.6615 19.1021 1.79337 19.5362 2.05706 19.5362H8.52214C8.6391 19.5362 8.74073 19.6154 8.77521 19.7271C8.99911 20.4529 9.99866 23 13 23C16.0013 23 17.0009 20.4529 17.2248 19.7271C17.2593 19.6154 17.3609 19.5362 17.4779 19.5362H23.9429C24.2066 19.5362 24.3385 19.1021 24.1359 18.9333C22.787 17.809 20.3333 14.9484 20.3333 9.01449Z"
              stroke="#333333"
              strokeWidth="2"
            />
          </Svg>
        </View>
      )}
      {showProfileView && (
        <View style={styles.profile}>
          <Svg width="24" height="25" viewBox="0 0 24 25" fill="none">
            <Circle
              cx="11.8877"
              cy="7.62015"
              r="3.61136"
              stroke="#333333"
              strokeWidth="2"
            />
            <Path
              d="M4.31738 24.0197V19.4482C4.31738 17.2391 6.10824 15.4482 8.31738 15.4482H16.1251C18.3342 15.4482 20.1251 17.2391 20.1251 19.4482V24.03"
              stroke="#333333"
              strokeWidth="2"
            />
          </Svg>
        </View>
      )}
      <View style={styles.search}>
        <Svg width="24" height="25" viewBox="0 0 24 25" fill="none">
          <Path
            d="M11 20.4844C15.9706 20.4844 20 16.4549 20 11.4844C20 6.51381 15.9706 2.48438 11 2.48438C6.02944 2.48438 2 6.51381 2 11.4844C2 16.4549 6.02944 20.4844 11 20.4844Z"
            stroke="#333333"
            strokeWidth="2"
          />
          <Path d="M22 22.4844L18 18.4844" stroke="#333333" strokeWidth="2" />
        </Svg>
      </View>

      <View style={styles.menu}>
        <Svg width="24" height="25" viewBox="0 0 24 25" fill="none">
          <Path d="M2 12.6924H22" stroke="#222222" strokeWidth="2" />
          <Path d="M2 4.69238H22" stroke="#222222" strokeWidth="2" />
          <Path d="M2 20.6924H22" stroke="#222222" strokeWidth="2" />
        </Svg>
      </View>
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
    padding: 10,
    position: "absolute",
    right: 50,
  },
  menu: {
    padding: 10,
    position: "absolute",
    right: 0,
  },
  notification: {
    padding: 10,
    position: "absolute",
    right: 90,
  },
  profile: {
    padding: 10,
    position: "absolute",
    right: 90,
  },
});

export default Head;

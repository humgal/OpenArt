import React from "react";
import { View, StyleSheet } from "react-native";
import Icon from "@expo/vector-icons/AntDesign";
import { IconButton } from "@react-native-material/core";
import * as RootNavigation from '../RootNavigation'

const BottomTab = () => {

  return (
    <View style={styles.footer}>
      <View style={{ flex: 2,alignItems:"center"}} >
      <IconButton  
      onPress={()=>RootNavigation.navigate('Home')}
      
      icon={props => <Icon name="home" {...props} />} />
      </View>
      <View style={{ flex: 2,alignItems:"center" }} >
      <IconButton 
      onPress={()=>RootNavigation.navigate('Creator')}
      icon={props => <Icon name="hearto" {...props} />} />
      </View> 
      <View style={{ flex: 2, alignItems:"center" }} >
      <IconButton 
      onPress={()=>RootNavigation.navigate('UploadArtWork')}
      icon={props => <Icon name="clouduploado" {...props} />} />
      </View>
      <View style={{ flex: 2, alignItems:"center" }} >
      <IconButton 
      onPress={()=>RootNavigation.navigate('Profile')}
      icon={props => <Icon name="profile" {...props} />} />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  footer: {
    position: "absolute",
    bottom: 0,
    width: "100%",
    height: 80,
    backgroundColor: "#ffffff",
    flexDirection: "row",
  },
});
export default BottomTab;

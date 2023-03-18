import React from 'react';
import { View ,Image,StyleSheet} from 'react-native';
import Svg,{Path} from 'react-native-svg';


const Head = ()=> {
    return(
        <View style={styles.container}>
            <View>
            <Image style={styles.tinyLogo}
            source={require('../../../../assets/Logo.png')}
            />
            </View>
            <View style={styles.search} >
                <Svg width="24" height="25" viewBox="0 0 24 25" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <Path d="M11 20.4844C15.9706 20.4844 20 16.4549 20 11.4844C20 6.51381 15.9706 2.48438 11 2.48438C6.02944 2.48438 2 6.51381 2 11.4844C2 16.4549 6.02944 20.4844 11 20.4844Z" stroke="#333333" stroke-width="2"/>
                    <Path d="M22 22.4844L18 18.4844" stroke="#333333" stroke-width="2"/>
                </Svg>
            </View>
           
            <View style={styles.menu}>
                <Svg width="24" height="25" viewBox="0 0 24 25" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <Path d="M2 12.6924H22" stroke="#222222" stroke-width="2"/>
                    <Path d="M2 4.69238H22" stroke="#222222" stroke-width="2"/>
                    <Path d="M2 20.6924H22" stroke="#222222" stroke-width="2"/>
                </Svg>
            </View>
            

        </View>
    )
};

const styles = StyleSheet.create({
    container: {
        paddingTop: 50,
        flexDirection: 'row',
        padding: 10,
        
      },
      tinyLogo: {
        width: 120,
        height: 30,
        
      },
     search: {
        paddingTop: 50,
        padding: 10,
        position: 'absolute',
        right:50
     },
     menu: {
        paddingTop: 50,
        padding: 10,
        position: 'absolute',
        right: 0,

     }
})

export default Head;
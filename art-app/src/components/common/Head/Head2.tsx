import React from 'react';
import { View ,Image,StyleSheet} from 'react-native';
import Svg,{Path} from 'react-native-svg';


const Head = ()=> {
    return(
        <View style={styles.container}>

            <Image style={styles.tinyLogo}
            source={require('../../../../assets/Logo.png')}/>
    
            <View style={styles.search} >
                <Svg width="24" height="25" viewBox="0 0 24 25" fill="none" >
                    <Path d="M11 20.4844C15.9706 20.4844 20 16.4549 20 11.4844C20 6.51381 15.9706 2.48438 11 2.48438C6.02944 2.48438 2 6.51381 2 11.4844C2 16.4549 6.02944 20.4844 11 20.4844Z" stroke="#333333" strokeWidth="2"/>
                    <Path d="M22 22.4844L18 18.4844" stroke="#333333" strokeWidth="2"/>
                </Svg>
            </View>
           
            <View style={styles.close}>
                <Svg width="24" height="25" viewBox="0 0 24 25" fill="none" >
                    <Path d="M6 6.69238L18.7742 19.4666" stroke="#14142B" strokeWidth="2" stroke-linejoin="round"/>
                    <Path d="M6 19.4668L18.7742 6.6926" stroke="#14142B" strokeWidth="2" stroke-linejoin="round"/>
                </Svg>

            </View>
            

        </View>
    )
};

const styles = StyleSheet.create({
    container: {
        marginTop: 50,
        flexDirection: 'row',
        padding: 10,
        
      },
      tinyLogo: {
        width: 120,
        height: 30,
        
      },
     search: {
       
        padding: 10,
        position: 'absolute',
        right:50
     },
     close: {
        
        padding: 10,
        position: 'absolute',
        right: 0,

     }
})

export default Head;
namespace java  com.ganji.arch.ip2city.finagleservice.thrift


struct  IpModel{  
      1:string sip ;  
      2:string eip;  
      3:string name;  
      4:i32 prov_id;  
      5:string prov_name;
      6:i32 city_id;
      7:string city_name;
      8:string city_pinyin;
      9:i32 dist_id;
      10:string  dist_name ;
                 }  





service  IpService {

  string getIpString(1:string ip)
  IpModel getIpModel(1:string ip)
   i32 blocking_call();

}

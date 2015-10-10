package add_record_infile

import (
	"domains"
	"os"
)

func Add(domaincsv domains.Domaincsv) {
 file, err := os.OpenFile("/home/juno/git/imagehoster_redis/usefull.bash", os.O_APPEND|os.O_WRONLY,0600)
     if err != nil {
         panic(err)
     }
     defer file.Close()
     
     str :="\nbin/contents_feeder_redis --site="+domaincsv.Name+" && sleep 90 && bin/contents_feeder_redis --site="+domaincsv.Name+" && sleep 90 && bin/contents_feeder_redis --site="+domaincsv.Name

     if _, err = file.WriteString(str); err != nil {
      panic(err)
     }
}

for D in `find src/ -type d`
do
    echo $D
    if  [[ $D == ./.git/* ]] 
    then
        continue    
    fi

    if  [[ $D == .. ]] 
    then
        continue    
    fi

    if  [[ $D == . ]] 
    then
        continue    
    fi

    cd $D
    pwd
    go test -coverprofile=cover.out
    gocov convert cover.out | gocov-xml > coverage.xml
    cd .. 
done
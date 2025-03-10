// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFrpUser []*NgingFrpUser

func (s Slice_NgingFrpUser) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpUser) RangeRaw(fn func(m *NgingFrpUser) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpUser) GroupBy(keyField string) map[string][]*NgingFrpUser {
	r := map[string][]*NgingFrpUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFrpUser{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFrpUser) KeyBy(keyField string) map[string]*NgingFrpUser {
	r := map[string]*NgingFrpUser{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFrpUser) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFrpUser) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFrpUser) FromList(data interface{}) Slice_NgingFrpUser {
	values, ok := data.([]*NgingFrpUser)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFrpUser{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingFrpUser(ctx echo.Context) *NgingFrpUser {
	m := &NgingFrpUser{}
	m.SetContext(ctx)
	return m
}

// NgingFrpUser 网络用户
type NgingFrpUser struct {
	base    factory.Base
	objects []*NgingFrpUser

	Id           uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	ServerId     uint   `db:"server_id" bson:"server_id" comment:"服务ID" json:"server_id" xml:"server_id"`
	CustomerId   uint64 `db:"customer_id" bson:"customer_id" comment:"客户ID" json:"customer_id" xml:"customer_id"`
	Uid          uint   `db:"uid" bson:"uid" comment:"管理员ID" json:"uid" xml:"uid"`
	Username     string `db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password     string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Banned       string `db:"banned" bson:"banned" comment:"是否禁止连接" json:"banned" xml:"banned"`
	Bandwidth    string `db:"bandwidth" bson:"bandwidth" comment:"带宽(MB/KB) " json:"bandwidth" xml:"bandwidth"`
	TrafficUsed  uint64 `db:"traffic_used" bson:"traffic_used" comment:"已用流量" json:"traffic_used" xml:"traffic_used"`
	TrafficTotal uint64 `db:"traffic_total" bson:"traffic_total" comment:"流量总量" json:"traffic_total" xml:"traffic_total"`
	IpWhitelist  string `db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个) " json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist  string `db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个) " json:"ip_blacklist" xml:"ip_blacklist"`
	Start        uint   `db:"start" bson:"start" comment:"生效时间" json:"start" xml:"start"`
	End          uint   `db:"end" bson:"end" comment:"过期时间" json:"end" xml:"end"`
	Created      uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated      uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingFrpUser) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingFrpUser) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFrpUser) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFrpUser) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFrpUser) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFrpUser) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFrpUser) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFrpUser) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingFrpUser) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFrpUser) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFrpUser) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFrpUser) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingFrpUser) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingFrpUser) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingFrpUser) Objects() []*NgingFrpUser {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFrpUser) XObjects() Slice_NgingFrpUser {
	return Slice_NgingFrpUser(a.Objects())
}

func (a *NgingFrpUser) NewObjects() factory.Ranger {
	return &Slice_NgingFrpUser{}
}

func (a *NgingFrpUser) InitObjects() *[]*NgingFrpUser {
	a.objects = []*NgingFrpUser{}
	return &a.objects
}

func (a *NgingFrpUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFrpUser) Short_() string {
	return "nging_frp_user"
}

func (a *NgingFrpUser) Struct_() string {
	return "NgingFrpUser"
}

func (a *NgingFrpUser) Name_() string {
	b := a
	if b == nil {
		b = &NgingFrpUser{}
	}
	if b.base.Namer() != nil {
		return WithPrefix(b.base.Namer()(b))
	}
	return WithPrefix(factory.TableNamerGet(b.Short_())(b))
}

func (a *NgingFrpUser) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFrpUser) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingFrpUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFrpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpUser(*v))
		case []*NgingFrpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpUser) GroupBy(keyField string, inputRows ...[]*NgingFrpUser) map[string][]*NgingFrpUser {
	var rows Slice_NgingFrpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpUser(inputRows[0])
	} else {
		rows = Slice_NgingFrpUser(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFrpUser) KeyBy(keyField string, inputRows ...[]*NgingFrpUser) map[string]*NgingFrpUser {
	var rows Slice_NgingFrpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpUser(inputRows[0])
	} else {
		rows = Slice_NgingFrpUser(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFrpUser) AsKV(keyField string, valueField string, inputRows ...[]*NgingFrpUser) param.Store {
	var rows Slice_NgingFrpUser
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpUser(inputRows[0])
	} else {
		rows = Slice_NgingFrpUser(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFrpUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFrpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpUser(*v))
		case []*NgingFrpUser:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpUser(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpUser) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFrpUser) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingFrpUser) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingFrpUser) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFrpUser) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFrpUser) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFrpUser) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFrpUser) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["banned"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["banned"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	editColumns := make([]string, 0, len(kvset))
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingFrpUser) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["banned"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["banned"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Updatex()
	}
	m := *a
	m.FromRow(kvset)
	editColumns := make([]string, 0, len(kvset))
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(kvset).Updatex(); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", &m, editColumns, mw, args...)
	return
}

func (a *NgingFrpUser) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingFrpUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingFrpUser) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingFrpUser) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingFrpUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFrpUser) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFrpUser) Reset() *NgingFrpUser {
	a.Id = 0
	a.ServerId = 0
	a.CustomerId = 0
	a.Uid = 0
	a.Username = ``
	a.Password = ``
	a.Banned = ``
	a.Bandwidth = ``
	a.TrafficUsed = 0
	a.TrafficTotal = 0
	a.IpWhitelist = ``
	a.IpBlacklist = ``
	a.Start = 0
	a.End = 0
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingFrpUser) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["ServerId"] = a.ServerId
		r["CustomerId"] = a.CustomerId
		r["Uid"] = a.Uid
		r["Username"] = a.Username
		r["Password"] = a.Password
		r["Banned"] = a.Banned
		r["Bandwidth"] = a.Bandwidth
		r["TrafficUsed"] = a.TrafficUsed
		r["TrafficTotal"] = a.TrafficTotal
		r["IpWhitelist"] = a.IpWhitelist
		r["IpBlacklist"] = a.IpBlacklist
		r["Start"] = a.Start
		r["End"] = a.End
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "ServerId":
			r["ServerId"] = a.ServerId
		case "CustomerId":
			r["CustomerId"] = a.CustomerId
		case "Uid":
			r["Uid"] = a.Uid
		case "Username":
			r["Username"] = a.Username
		case "Password":
			r["Password"] = a.Password
		case "Banned":
			r["Banned"] = a.Banned
		case "Bandwidth":
			r["Bandwidth"] = a.Bandwidth
		case "TrafficUsed":
			r["TrafficUsed"] = a.TrafficUsed
		case "TrafficTotal":
			r["TrafficTotal"] = a.TrafficTotal
		case "IpWhitelist":
			r["IpWhitelist"] = a.IpWhitelist
		case "IpBlacklist":
			r["IpBlacklist"] = a.IpBlacklist
		case "Start":
			r["Start"] = a.Start
		case "End":
			r["End"] = a.End
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpUser) FromRow(row map[string]interface{}) {
	for key, value := range row {
		if _, ok := value.(db.RawValue); ok {
			continue
		}
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "server_id":
			a.ServerId = param.AsUint(value)
		case "customer_id":
			a.CustomerId = param.AsUint64(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "username":
			a.Username = param.AsString(value)
		case "password":
			a.Password = param.AsString(value)
		case "banned":
			a.Banned = param.AsString(value)
		case "bandwidth":
			a.Bandwidth = param.AsString(value)
		case "traffic_used":
			a.TrafficUsed = param.AsUint64(value)
		case "traffic_total":
			a.TrafficTotal = param.AsUint64(value)
		case "ip_whitelist":
			a.IpWhitelist = param.AsString(value)
		case "ip_blacklist":
			a.IpBlacklist = param.AsString(value)
		case "start":
			a.Start = param.AsUint(value)
		case "end":
			a.End = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingFrpUser) GetField(field string) interface{} {
	switch field {
	case "Id":
		return a.Id
	case "ServerId":
		return a.ServerId
	case "CustomerId":
		return a.CustomerId
	case "Uid":
		return a.Uid
	case "Username":
		return a.Username
	case "Password":
		return a.Password
	case "Banned":
		return a.Banned
	case "Bandwidth":
		return a.Bandwidth
	case "TrafficUsed":
		return a.TrafficUsed
	case "TrafficTotal":
		return a.TrafficTotal
	case "IpWhitelist":
		return a.IpWhitelist
	case "IpBlacklist":
		return a.IpBlacklist
	case "Start":
		return a.Start
	case "End":
		return a.End
	case "Created":
		return a.Created
	case "Updated":
		return a.Updated
	default:
		return nil
	}
}

func (a *NgingFrpUser) GetAllFieldNames() []string {
	return []string{
		"Id",
		"ServerId",
		"CustomerId",
		"Uid",
		"Username",
		"Password",
		"Banned",
		"Bandwidth",
		"TrafficUsed",
		"TrafficTotal",
		"IpWhitelist",
		"IpBlacklist",
		"Start",
		"End",
		"Created",
		"Updated",
	}
}

func (a *NgingFrpUser) HasField(field string) bool {
	switch field {
	case "Id":
		return true
	case "ServerId":
		return true
	case "CustomerId":
		return true
	case "Uid":
		return true
	case "Username":
		return true
	case "Password":
		return true
	case "Banned":
		return true
	case "Bandwidth":
		return true
	case "TrafficUsed":
		return true
	case "TrafficTotal":
		return true
	case "IpWhitelist":
		return true
	case "IpBlacklist":
		return true
	case "Start":
		return true
	case "End":
		return true
	case "Created":
		return true
	case "Updated":
		return true
	default:
		return false
	}
}

func (a *NgingFrpUser) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint64(vv)
		case "ServerId":
			a.ServerId = param.AsUint(vv)
		case "CustomerId":
			a.CustomerId = param.AsUint64(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "Username":
			a.Username = param.AsString(vv)
		case "Password":
			a.Password = param.AsString(vv)
		case "Banned":
			a.Banned = param.AsString(vv)
		case "Bandwidth":
			a.Bandwidth = param.AsString(vv)
		case "TrafficUsed":
			a.TrafficUsed = param.AsUint64(vv)
		case "TrafficTotal":
			a.TrafficTotal = param.AsUint64(vv)
		case "IpWhitelist":
			a.IpWhitelist = param.AsString(vv)
		case "IpBlacklist":
			a.IpBlacklist = param.AsString(vv)
		case "Start":
			a.Start = param.AsUint(vv)
		case "End":
			a.End = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingFrpUser) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["server_id"] = a.ServerId
		r["customer_id"] = a.CustomerId
		r["uid"] = a.Uid
		r["username"] = a.Username
		r["password"] = a.Password
		r["banned"] = a.Banned
		r["bandwidth"] = a.Bandwidth
		r["traffic_used"] = a.TrafficUsed
		r["traffic_total"] = a.TrafficTotal
		r["ip_whitelist"] = a.IpWhitelist
		r["ip_blacklist"] = a.IpBlacklist
		r["start"] = a.Start
		r["end"] = a.End
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "server_id":
			r["server_id"] = a.ServerId
		case "customer_id":
			r["customer_id"] = a.CustomerId
		case "uid":
			r["uid"] = a.Uid
		case "username":
			r["username"] = a.Username
		case "password":
			r["password"] = a.Password
		case "banned":
			r["banned"] = a.Banned
		case "bandwidth":
			r["bandwidth"] = a.Bandwidth
		case "traffic_used":
			r["traffic_used"] = a.TrafficUsed
		case "traffic_total":
			r["traffic_total"] = a.TrafficTotal
		case "ip_whitelist":
			r["ip_whitelist"] = a.IpWhitelist
		case "ip_blacklist":
			r["ip_blacklist"] = a.IpBlacklist
		case "start":
			r["start"] = a.Start
		case "end":
			r["end"] = a.End
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpUser) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPage(a, cond, sorts...)
}

func (a *NgingFrpUser) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageAs(a, recv, cond, sorts...)
}

func (a *NgingFrpUser) ListPageByOffset(cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageByOffset(a, cond, sorts...)
}

func (a *NgingFrpUser) ListPageByOffsetAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageByOffsetAs(a, recv, cond, sorts...)
}

func (a *NgingFrpUser) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFrpUser) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}

// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package product

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Product struct {
	Id          int64   `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Name        string  `thrift:"name,2" frugal:"2,default,string" json:"name"`
	Price       *int64  `thrift:"price,3,optional" frugal:"3,optional,i64" json:"price,omitempty"`
	Stock       *int64  `thrift:"stock,4,optional" frugal:"4,optional,i64" json:"stock,omitempty"`
	Image       *string `thrift:"image,5,optional" frugal:"5,optional,string" json:"image,omitempty"`
	Description *string `thrift:"description,6,optional" frugal:"6,optional,string" json:"description,omitempty"`
	CreateTime  *string `thrift:"create_time,7,optional" frugal:"7,optional,string" json:"create_time,omitempty"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) InitDefault() {
	*p = Product{}
}

func (p *Product) GetId() (v int64) {
	return p.Id
}

func (p *Product) GetName() (v string) {
	return p.Name
}

var Product_Price_DEFAULT int64

func (p *Product) GetPrice() (v int64) {
	if !p.IsSetPrice() {
		return Product_Price_DEFAULT
	}
	return *p.Price
}

var Product_Stock_DEFAULT int64

func (p *Product) GetStock() (v int64) {
	if !p.IsSetStock() {
		return Product_Stock_DEFAULT
	}
	return *p.Stock
}

var Product_Image_DEFAULT string

func (p *Product) GetImage() (v string) {
	if !p.IsSetImage() {
		return Product_Image_DEFAULT
	}
	return *p.Image
}

var Product_Description_DEFAULT string

func (p *Product) GetDescription() (v string) {
	if !p.IsSetDescription() {
		return Product_Description_DEFAULT
	}
	return *p.Description
}

var Product_CreateTime_DEFAULT string

func (p *Product) GetCreateTime() (v string) {
	if !p.IsSetCreateTime() {
		return Product_CreateTime_DEFAULT
	}
	return *p.CreateTime
}
func (p *Product) SetId(val int64) {
	p.Id = val
}
func (p *Product) SetName(val string) {
	p.Name = val
}
func (p *Product) SetPrice(val *int64) {
	p.Price = val
}
func (p *Product) SetStock(val *int64) {
	p.Stock = val
}
func (p *Product) SetImage(val *string) {
	p.Image = val
}
func (p *Product) SetDescription(val *string) {
	p.Description = val
}
func (p *Product) SetCreateTime(val *string) {
	p.CreateTime = val
}

var fieldIDToName_Product = map[int16]string{
	1: "id",
	2: "name",
	3: "price",
	4: "stock",
	5: "image",
	6: "description",
	7: "create_time",
}

func (p *Product) IsSetPrice() bool {
	return p.Price != nil
}

func (p *Product) IsSetStock() bool {
	return p.Stock != nil
}

func (p *Product) IsSetImage() bool {
	return p.Image != nil
}

func (p *Product) IsSetDescription() bool {
	return p.Description != nil
}

func (p *Product) IsSetCreateTime() bool {
	return p.CreateTime != nil
}

func (p *Product) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField7(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Product[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Product) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *Product) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Name = _field
	return nil
}
func (p *Product) ReadField3(iprot thrift.TProtocol) error {

	var _field *int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Price = _field
	return nil
}
func (p *Product) ReadField4(iprot thrift.TProtocol) error {

	var _field *int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Stock = _field
	return nil
}
func (p *Product) ReadField5(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Image = _field
	return nil
}
func (p *Product) ReadField6(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Description = _field
	return nil
}
func (p *Product) ReadField7(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.CreateTime = _field
	return nil
}

func (p *Product) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Product"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Product) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Product) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Product) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetPrice() {
		if err = oprot.WriteFieldBegin("price", thrift.I64, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.Price); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Product) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetStock() {
		if err = oprot.WriteFieldBegin("stock", thrift.I64, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.Stock); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *Product) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetImage() {
		if err = oprot.WriteFieldBegin("image", thrift.STRING, 5); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Image); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *Product) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetDescription() {
		if err = oprot.WriteFieldBegin("description", thrift.STRING, 6); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Description); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *Product) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetCreateTime() {
		if err = oprot.WriteFieldBegin("create_time", thrift.STRING, 7); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.CreateTime); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}

func (p *Product) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Product(%+v)", *p)

}

func (p *Product) DeepEqual(ano *Product) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Name) {
		return false
	}
	if !p.Field3DeepEqual(ano.Price) {
		return false
	}
	if !p.Field4DeepEqual(ano.Stock) {
		return false
	}
	if !p.Field5DeepEqual(ano.Image) {
		return false
	}
	if !p.Field6DeepEqual(ano.Description) {
		return false
	}
	if !p.Field7DeepEqual(ano.CreateTime) {
		return false
	}
	return true
}

func (p *Product) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *Product) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Name, src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field3DeepEqual(src *int64) bool {

	if p.Price == src {
		return true
	} else if p.Price == nil || src == nil {
		return false
	}
	if *p.Price != *src {
		return false
	}
	return true
}
func (p *Product) Field4DeepEqual(src *int64) bool {

	if p.Stock == src {
		return true
	} else if p.Stock == nil || src == nil {
		return false
	}
	if *p.Stock != *src {
		return false
	}
	return true
}
func (p *Product) Field5DeepEqual(src *string) bool {

	if p.Image == src {
		return true
	} else if p.Image == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Image, *src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field6DeepEqual(src *string) bool {

	if p.Description == src {
		return true
	} else if p.Description == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Description, *src) != 0 {
		return false
	}
	return true
}
func (p *Product) Field7DeepEqual(src *string) bool {

	if p.CreateTime == src {
		return true
	} else if p.CreateTime == nil || src == nil {
		return false
	}
	if strings.Compare(*p.CreateTime, *src) != 0 {
		return false
	}
	return true
}
////
////  CameraButtonsView.swift
////  PINATA
////
////  Created by Karim Abedrabbo on 11/19/19.
////  Copyright © 2019 Karim Abedrabbo. All rights reserved.
////
//
//import Foundation
//import UIKit
//
//public class CameraButtonsView: UIView {
//  internal var setToDelete = false
//
//  var flipButton: UIButton = {
//      let flip = UIButton(type: .custom)
//      flip.setImage(UIImage(named: "flip_button"), for: .normal)
//      flip.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
//
//      flip.addTarget(self, action: #selector(handleFlipButton(_:)), for: .touchUpInside)
//      flip.isHidden = false
//      flip.translatesAutoresizingMaskIntoConstraints = false
//      return flip
//  }()
//  
//  var flashButton: UIButton = {
//      let flash = UIButton(type: .custom)
//      flash.setImage(UIImage(named: "flash_off_button"), for: .normal)
//      flash.setImage(UIImage(named: "flash_on_button"), for: .selected)
//      flash.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
//      flash.addTarget(self, action: #selector(handleFlashModeButton(_:)), for: .touchUpInside)
//      flash.isHidden = false
//      flash.translatesAutoresizingMaskIntoConstraints = false
//      return flash
//  }()
//  
//  var mediaButton: UIButton = {
//      let media = UIButton(type: .custom)
//      media.setImage(UIImage(named: "media_button"), for: .normal)
//      media.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
//
//      media.addTarget(self, action: #selector(handleMediaButton(_:)), for: .touchUpInside)
//      media.isHidden = false
//      media.translatesAutoresizingMaskIntoConstraints = false
//      return media
//  }()
//  
//
//  
//  //doesnt work as a view but works as a button WITH NO TARGET TAP HANDLED BY GESTURES because there is also a long hold gesture
//  var deleteSegmentView: UIButton = {
//
//      let deleteSegment = UIButton(type: .custom)
//      deleteSegment.setImage( UIImage(named: "delete_segment_button"), for: .normal)
//
//    deleteSegment.contentMode = .center
//      deleteSegment.isHidden = true
//      deleteSegment.translatesAutoresizingMaskIntoConstraints = false
//      
//      deleteSegment.layer.cornerRadius = 30
//      return deleteSegment
//      
//  }()
//  
//  var cancelButton: UIButton = {
//      let cancel = UIButton(type: .custom)
//      cancel.setImage(UIImage(named: "cancel_button"), for: .normal)
//      cancel.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
//
//      cancel.addTarget(self, action: #selector(handleCancelButton(_:)), for: .touchUpInside)
//      cancel.isHidden = false
//      cancel.translatesAutoresizingMaskIntoConstraints = false
//      return cancel
//  }()
//  
//  
//  var doneRecordingButton: UIButton = {
//      let doneRecording = UIButton(type: .custom)
//      var completeImage = UIImage(named: "complete_button")
//
//      doneRecording.setImage(completeImage, for: .normal)
//
//    doneRecording.setTitle("  Next", for: .normal)
//      doneRecording.titleLabel?.font = doneRecording.titleLabel?.font.withSize(20)
//      doneRecording.isEnabled = false
//      doneRecording.addTarget(self, action: #selector(handleDoneRecordingButton(_:)), for: .touchUpInside)
//      doneRecording.backgroundColor = UIColor(red: 210.0/255.0, green: 40/255.0, blue: 40.0/255.0, alpha: 1.0)
//      doneRecording.titleLabel?.font = UIFont.boldSystemFont(ofSize: 16.0)
//      
//      doneRecording.layer.cornerRadius = 30.0
//      doneRecording.isHidden = true
//      doneRecording.translatesAutoresizingMaskIntoConstraints = false
//      return doneRecording
//  }()
//}
